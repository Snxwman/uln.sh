package ln

import (
	"database/sql"
	"math/rand"
	"strings"
)

const DEBUG bool = true

var BASE_URL string
var lnApp LnApp

var chars = []string {
    "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", 
    "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
    "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
    "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
    "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

type LnApp struct {
    // TODO: Per service config
    db *sql.DB
    urls map[string]*shortlink
    urlsReverse map[string]string
}

type PathExistsError struct {}
func (e PathExistsError) Error() string {
   return "Path already exists" 
}

func Init(db *sql.DB) {
    if DEBUG {
        BASE_URL = "http://localhost:8080"
    } else {
        BASE_URL = "https://uln.sh"
    }
    
    lnApp = LnApp {
        db: nil,
        urls: make(map[string]*shortlink),
        urlsReverse: make(map[string]string),
    }

    
    // initDatabase()
}

func makePath(length int) string {
    var path string

    if length == 0 {
        length = 7
    }

    for i := 0; i < length; i++ {
        path += chars[rand.Intn(len(chars))]
    }

    return path
}

func pathExists(path string) bool {
    _, ok := lnApp.urls[path]
    if !ok {
        return false
    }
    return true
}

func getNextPath() string {
    return ""
}

func shortlinkExists(rawURL string) (string, bool) {
    path, exists := lnApp.urlsReverse[rawURL]
    if exists {
        return path, true
    } else {
        return "", false
    }
}

func registerShortlink(s *shortlink) error {
    path := strings.Trim(s.shortURL.Path, "/")
    if pathExists(path) {
        return PathExistsError{}
    } else {
        lnApp.urls[path] = s
        lnApp.urlsReverse[s.fullURL.String()] = path
        return nil
    }
}
