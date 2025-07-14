package store

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx"
    _ "github.com/jackc/pgx/v5/stdlib"
)

type QueryString string

const DefaultDbTimeout time.Duration = 5*time.Second

const createCreatedViaEnumQuery QueryString = `

    CREATE TYPE created_via AS ENUM (
        'cli',
        'uln-cli',
        'uln-tui',
        'web',
        'mobile'
    );
`

const createCreationEventsTableQuery QueryString = `
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

const (
    host = "uln-postgres"
    port = 5432
    user = "postgres"
    password = "secret"
    dbname = "uln"
)

func Init() *sqlx.DB {
    postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", 
        host, port, user, password, dbname)

    db := sqlx.MustConnect("pgx", postgresInfo)

    initTables(db)

    return db
}

func initTables(db *sqlx.DB) error {
    db.Exec(string(createCreatedViaEnumQuery))
    db.Exec(string(createCreationEventsTableQuery))

    return nil
}
