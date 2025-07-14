package ln

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/jmoiron/sqlx"
)

const DEBUG bool = true

var BASE_URL string

var chars = []string {
    "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", 
    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
    "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
    "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
    "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

var lnApp LnApp
type LnApp struct {
    // TODO: Per service config
    db lnDB
    cache lnCache
}

type PathExistsError struct {}
func (e PathExistsError) Error() string {
   return "Path already exists" 
}

func Init(db *sqlx.DB) {
    if DEBUG {
        BASE_URL = "http://localhost:8080"
    } else {
        BASE_URL = "https://uln.sh"
    }
    
    lnApp = LnApp {
        db: newLnDB(db),
        cache: newLnCache(),
    }

    lnApp.db.initDatabase()
    lnApp.db.initTables()    
}

func makeRandomPath(length int) string {
    var path string

    if length == 0 {
        length = 7
    }

    for i := 0; i < length; i++ {
        path += chars[rand.Intn(len(chars))]
    }

    return path
}

func makeBase62Path(url string, length int) string {
    return ""
}

func validateCustomPath(path string) bool {
    // Check user perms to make a custom url
    // Check banned/reserved paths
    // Check cache and db for duplicates
    return false
}

func getNextPath() string {
    return ""
}

func tryGetShortlinkByPath(path string) (*shortlink, bool) {
    shortlinkFromCache, inCache := lnApp.cache.contains(path)
    shortlinkFromDb, inDb := lnApp.db.getShortlinkByPath(path)

    if !inCache && inDb != nil {
        return nil, false
    }

    if inCache {
        return shortlinkFromCache, true 
    } else {
        return shortlinkFromDb, true 
    }
}

func tryGetShortlinkForUrl(rawURL string) ([]*shortlink, bool) {
    shortlinkFromCache, inCache := lnApp.cache.containsReverse(rawURL)
    shortlinksFromDb, inDb := lnApp.db.getShortlinksForUrl(rawURL)
    fmt.Printf("URL in cache: %t \nURL in database: %t\n", inCache, inDb)

    if inCache {
        return []*shortlink{shortlinkFromCache}, true 
    } else if inDb {
        return shortlinksFromDb, true 
    }

    return nil, false
}

func registerShortlink(s *shortlink) error {
    // FIXME: Inappropriate place to do this check
    path := strings.Trim(s.ShortURL.Path, "/")
    if _, inCache := lnApp.cache.contains(path); inCache {
        return PathExistsError{}
    } else {
        lnApp.cache.insert(s)
        return nil
    }
}
