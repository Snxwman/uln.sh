package ln

import (
	"context"
	"log"
	"time"
    
    _ "github.com/lib/pq"
)

const defaultTimeout time.Duration = 5*time.Second

func defaultTableValues(s *shortlink) bool {
    zeroTime := time.Time{}

    if !s.active || 
        s.reserved || 
        s.redirectReqs != 0 || 
        s.infoReqs != 0 || 
        s.lastAccessed != zeroTime {
        return false
    }

    return true
}

func initTables() error {
    query := `
    CREATE TABLE IF NOT EXISTS 
    shortlinks(
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

    ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
    defer cancel()

    _, err := lnApp.db.ExecContext(ctx, query)
    if err != nil {
        // TODO: Customize error messages
        log.Printf("%s", err)
        return err
    }

    log.Printf("")

    return nil
}

func destroyDatabase() error {
    return nil
}

func getShortlinkByPath(path string) error {
    return nil
}

func getShortlinkByID(id string) error {
    return nil
}

// func insertShortlink(s *shortlink) error {
//     query := `
//     INSERT INTO shortlinks (full_url, short_url, expiration)
//     VALUES($1, $2, $3);
//     `
//
//     queryParams := queryParams {
//         query: query,
//         args: [3]any{s.fullURL, s.shortURL, s.expiration},
//         successMessage: "Inserted shortlink into table",
//         errorMessage: "Failed to insert shortlink",
//         timeout: 5*time.Second,
//     }
//
//     err := execQuery(queryParams)
//     if err != nil {
//         return err
//     }
//
//     if !defaultTableValues(s) {
//         err := updateShortlink(s, db)
//         if err != nil {
//             return err
//         }
//     }
//
//     return nil
// }

// func updateShortlink(s *shortlink) error {
//     query := `
//     ;
//     `
//
//     queryParams := queryParams {
//         query: query,
//         successMessage: "Updated row in shortlinks table",
//         errorMessage: "Failed to insert shortlink",
//         timeout: 5*time.Second,
//     }
//
//     err := execQuery(queryParams)
//     if err != nil {
//         return err
//     }
//     
//     return nil
// }

// func deleteShortlink(s *shortlink) error {
//     query := `
//     ;
//     `
//
//     queryParams := queryParams {
//         query: query,
//         successMessage: "Deleted row from shortlinks table",
//         errorMessage: "Failed to delete row from shortlinks table",
//         timeout: 5*time.Second,
//     }
//
//     err := execQuery(queryParams)
//     if err != nil {
//         return err
//     }
//
//     return nil
// }
