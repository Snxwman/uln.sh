package ln

import (
	"context"
	"database/sql"
	"time"
	"uln/src/store"

	_ "github.com/lib/pq"
)

const createShortlinkTypeEnumQuery store.QueryString = `
    CREATE TYPE shortlink_type AS ENUM (
        'random',
        'base62',
        'custom'
    );
`

const createShortlinkTableQuery store.QueryString = `
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

const createShortlinkCreationOptionsTableQuery store.QueryString = `
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

const insertShortlinkWithDefaultsQuery store.QueryString = `
    INSERT INTO 
    shortlinks (
        id,
        full_url, 
        short_url, 
        expiration
    )
    VALUES(gen_random_uuid(), $1, $2, $3);
`

const insertShortlinkFullQuery store.QueryString = `
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

const insertShortlinkCreationOptionsQuery store.QueryString = `
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

const getShortlinkByPathQuery store.QueryString = `
    SELECT * FROM shortlinks
    WHERE short_url = $1;
`

const getShortlinkByIdQuery store.QueryString = `
    SELECT * FROM shortlinks
    WHERE id = $1;
`

const getShortlinksForUserQuery store.QueryString = `

`

const getShortlinksForTokenQuery store.QueryString = `

`

const getShortlinksForIpQuery store.QueryString = `

`

const updateShortlinkQuery store.QueryString = `

`

const deleteShortlinkQuery store.QueryString = `

`

const deleteShortlinkCreationOptionsQuery store.QueryString = `

`

const defaultTimeout time.Duration = 5*time.Second

type queryParams struct {
    query store.QueryString
    timeout time.Duration
    errMsg string
    okMsg string
}

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

func execWithParams(params queryParams) error {
    ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
    defer cancel()
    
    _, err := lnApp.db.ExecContext(ctx, string(params.query))
    if err != nil {
        return err
    }
    return nil
}

func queryWithParams(params queryParams) (*sql.Rows, error) {
    ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
    defer cancel()
    
    rows, err := lnApp.db.QueryContext(ctx, string(params.query))
    if err != nil {
        return nil, err
    }

    return rows, nil
}

func initDatabase() error {
    err := execWithParams(queryParams {
        query: createShortlinkTypeEnumQuery,
        timeout: defaultTimeout,
        errMsg: "",
        okMsg: "",
    })

    if err != nil {
        return err
    }

    return nil
}

func initTables() error {
    params := queryParams {
        query: createShortlinkTableQuery,
        timeout: defaultTimeout,
        errMsg: "",
        okMsg: "",
    }

    err := execWithParams(params)
    if err != nil {
        return err
    }

    params.query = createShortlinkCreationOptionsTableQuery
    err = execWithParams(params)
    if err != nil {
        return err
    }

    return nil
}

func destroyDatabase() error {
    return nil
}

func getShortlinkByPath(path string) (shortlink, error) {
    return shortlink{}, nil
}

func getShortlinkByID(id string) (shortlink, error) {
    return shortlink{}, nil
}

func getShortlinksForUser() ([]shortlink, error) {
    return []shortlink{}, nil
}

func getShortlinksForToken() ([]shortlink, error) {
    return []shortlink{}, nil
}

func getShortlinksForIp() ([]shortlink, error) {
    return []shortlink{}, nil
}

func insertShortlink(s *shortlink) error {
    return nil
}

func insertShortlinkCreationOptions() error {
    return nil
}

func updateShortlink(s *shortlink) error {
    return nil
}

func deleteShortlink(s *shortlink) error {
    return nil
}

func deleteShortlinkCreationOptions() error {
    return nil
}
