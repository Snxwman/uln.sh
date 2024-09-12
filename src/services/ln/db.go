package ln

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func initDatabase(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS 
    shortlinks(
        id int primary key auto_increment
        full_url string
        short_url string
        active bool default true
        reserved bool default false
        expiration datetime
        redirect_reqs int default 0
        info_reqs int default 0
        last_accessed datetime
    )
    `
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := db.ExecContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when creating shortlinks table", err)
        return err
    }

    log.Printf("Shortlinks table initialized")
    return nil
}

func insertShortlink(s *shortlink, db *sql.DB) error {
    return nil
}

func updateShortlink(s *shortlink, db *sql.DB) error {
    return nil
}

func deleteShortlink(s *shortlink, db *sql.DB) error {
    return nil
}
