package ln

import (
	"fmt"
	"log"
    "net/url"
	"time"

	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"

	"uln/src/store"
)

type lnDB struct {
    *sqlx.DB
    queries lnQueries
}

func newLnDB(db *sqlx.DB) lnDB {
    lndb := new(lnDB)
    lndb.DB = db
    lndb.queries.init()
    return *lndb
}

type lnQueries struct {
    create struct {
        enum struct {
            shortlinkType store.QueryString
        }
        table struct {
            shortlink store.QueryString
            shortlinkCreationOptions store.QueryString
        }
    }
    insert struct {
        shortlink struct {
            withDefaults store.QueryString
            withAllColumns store.QueryString
        }
        shortlinkCreationOptions store.QueryString
    }
    get struct {
        shortlink struct {
            byPath store.QueryString
            byId store.QueryString
            forFullURL store.QueryString
            forUser store.QueryString
            forToken store.QueryString
            forIp store.QueryString
        }
        shortlinkCreationOptions struct {
            byPath store.QueryString
            byId store.QueryString
        }
    }
    update struct {
        shortlink store.QueryString
    }
    delete struct {
        shortlink store.QueryString
        shortlinkCreationOptions store.QueryString
    }
}

func (q *lnQueries) init() {
    q.create.enum.shortlinkType = `
        CREATE TYPE shortlink_type AS ENUM (
            'random',
            'base62',
            'custom'
        );
    `
    q.create.table.shortlink = `
        CREATE TABLE IF NOT EXISTS 
        shortlinks (
            id UUID PRIMARY KEY,
            full_url TEXT NOT NULL,
            short_url TEXT UNIQUE NOT NULL,
            active BOOL DEFAULT true,
            reserved BOOL DEFAULT false,
            redirect_reqs INT8 DEFAULT 0,
            info_reqs INT8 DEFAULT 0,
            expiration TIMESTAMP,
            last_accessed TIMESTAMP DEFAULT null
        );
    `
    q.create.table.shortlinkCreationOptions = `
        CREATE TABLE IF NOT EXISTS
        shortlink_creation_options (
            id UUID PRIMARY KEY,
            shortlink_id UUID,
            creation_event_id UUID,
            shortlink_type shortlink_type,
            keep_unique BOOL DEFAULT true,
            CONSTRAINT fk_shortlink_id
                FOREIGN KEY(shortlink_id)
                    REFERENCES shortlinks(id),
            CONSTRAINT fk_creation_event
                FOREIGN KEY(creation_event_id)
                    REFERENCES creation_events(id)
        );
    `

    q.insert.shortlink.withDefaults = `
        INSERT INTO 
        shortlinks (
            id,
            full_url, 
            short_url, 
            expiration
        )
        VALUES(gen_random_uuid(), $1, $2, $3);
    `

    q.insert.shortlink.withAllColumns = `
        INSERT INTO
        shortlinks (
            id,
            full_url,
            short_url,
            active,
            reserved,
            redirect_reqs,
            info_reqs,
            expiration,
            last_accessed
        )
        VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7, $8);
    `

    q.insert.shortlinkCreationOptions = `
        INSERT INTO
        shortlink_creation_options (
            id,
            shortlink_id,
            creation_event_id,
            shortlink_type,
            keep_unique,
        )
        VALUES(gen_random_uuid(), $1, $2, $3, $4);
    `

    q.get.shortlink.byPath = `
        SELECT * FROM shortlinks
        WHERE short_url = $1;
    `

    q.get.shortlink.byId = `
        SELECT * FROM shortlinks
        WHERE id = $1;
    `

    q.get.shortlink.forFullURL = `
        SELECT 
            full_url, 
            short_url, 
            active, 
            reserved, 
            expiration, 
            redirect_reqs,
            info_reqs,
            last_accessed
        FROM shortlinks
        WHERE full_url LIKE '%' || $1 || '%';
    `

    q.get.shortlink.forUser = `

    `

    q.get.shortlink.forToken = `

    `

    q.get.shortlink.forIp = `

    `

    q.get.shortlinkCreationOptions.byPath = `

    `

    q.get.shortlinkCreationOptions.byId = `
        SELECT * FROM shortlink_creation_options
        WHERE id = $1;
    `

    q.update.shortlink = `
        UPDATE shortlinks
        SET
            active = $2
            reserved = $3
            redirect_reqs = $4
            info_reqs = $5
            expiration = $6
            last_accessed = $7
        WHERE
            id = $1;
    `

    q.delete.shortlink = `
        DELETE FROM shortlinks
        WHERE id = $1;
    `

    q.delete.shortlinkCreationOptions = `
        DELETE FROM shortlink_creation_options
        WHERE id = $1;
    `
}

// FIX: Might not need
func defaultTableValues(s *shortlink) bool {
    zeroTime := time.Time{}

    if !s.Active || 
        s.Reserved || 
        s.RedirectReqs != 0 || 
        s.InfoReqs != 0 || 
        s.LastAccessed != zeroTime {
        return false
    }

    return true
}

// TODO: URGENT
func (db *lnDB) initDatabase() error {
    return nil
}

// TODO: URGENT
func (db *lnDB) initTables() error {
    return nil
}

func (db *lnDB) destroyDatabase() error {
    return nil
}

func (db *lnDB) getShortlinkByPath(path string) (*shortlink, error) {
    var s shortlink
    err := lnApp.db.Get(&s, string(lnApp.db.queries.get.shortlink.byPath))
    if err != nil {
        return nil, err
    }

    return &s, nil
}

func (db *lnDB) getShortlinkByID(id string) (*shortlink, error) {
    return &shortlink{}, nil
}

func (db *lnDB) getShortlinksForUrl(u string) ([]*shortlink, bool) {
    rows, err := lnApp.db.Queryx(
        string(lnApp.db.queries.get.shortlink.forFullURL),
        u,
    )

    if err != nil {
        fmt.Println("%w", err)
        return nil, false
    }

    shortlinks := []*shortlink{}
    for rows.Next() {
        var shortlink shortlink
        var fullURL string
        var shortURL string
        rows.Scan(
            &fullURL,
            &shortURL,
            &shortlink.Active,
            &shortlink.Reserved,
            &shortlink.Expiration,
            &shortlink.RedirectReqs,
            &shortlink.InfoReqs,
            &shortlink.LastAccessed,
        )

        parsed, err := url.Parse(fullURL)
        if err != nil {
            break
        }
        shortlink.FullURL = *parsed

        parsed, err = url.Parse(shortURL)
        if err != nil {
            break
        }
        shortlink.ShortURL = *parsed
        
        shortlinks = append(shortlinks, &shortlink)
        fmt.Printf("%v\n", shortlink)
    } 

    fmt.Printf("db.go: %v\n", shortlinks)

    if len(shortlinks) == 0 {
        return nil, false
    } else {
        return shortlinks, true
    }
}

func (db *lnDB) getShortlinksForUser() ([]*shortlink, error) {
    return []*shortlink{}, nil
}

func (db *lnDB) getShortlinksForToken() ([]*shortlink, error) {
    return []*shortlink{}, nil
}

func (db *lnDB) getShortlinksForIp() ([]*shortlink, error) {
    return []*shortlink{}, nil
}

// FIX: eliminate execWithParams usage
func (db *lnDB) insertShortlink(s *shortlink) error {
    var err error
    if defaultTableValues(s) {
        lnApp.db.Exec(
            string(lnApp.db.queries.insert.shortlink.withDefaults),
            s.FullURL.String(),
            s.ShortURL.String(),
            s.Expiration,
        )
    } else {
        lnApp.db.Exec(
            string(lnApp.db.queries.insert.shortlink.withAllColumns),
            s.FullURL.String(), 
            s.ShortURL.String(), 
            s.Active, 
            s.Reserved, 
            s.RedirectReqs, 
            s.InfoReqs, 
            s.Expiration, 
            s.LastAccessed,
        )
    }

    if err != nil {
        log.Printf("Error: %s", err.Error())
        return err
    }

    return nil
}

func (db *lnDB) insertShortlinkCreationOptions() error {
    return nil
}

func (db *lnDB) updateShortlink(s *shortlink) error {
    return nil
}

func (db *lnDB) deleteShortlink(s *shortlink) error {
    return nil
}

func (db *lnDB) deleteShortlinkCreationOptions() error {
    return nil
}
