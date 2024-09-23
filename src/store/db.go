package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
    host = "uln-postgres"
    port = 5432
    user = "postgres"
    password = "secret"
    dbname = "uln"
)

func Init() *sql.DB {
    postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", postgresInfo)
    if err != nil {
        fmt.Println(err)
    }

    initTables(db)

    return db
}

func initTables(db *sql.DB) error {
    query := `
    CREATE TYPE created_via AS ENUM (
        'cli',    
        'web',
        'mobile'
    );

    CREATE TABLE IF NOT EXISTS 
    creation_events (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP,
        created_by_user TEXT,
        created_by_ip INET,
        created_via created_via,
        initial_creation BOOL DEFAULT true
    );
    `

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    _, err := db.ExecContext(ctx, query)
    if err != nil {
        return err
    }

    return nil
}
