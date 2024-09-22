package store

import (
	"database/sql"
	"fmt"
)

func Init() *sql.DB {
    username := "postgres"
    password := "password"
    host := "localhost"
    connURL := fmt.Sprintf("postgres://%s:%s@%s", username, password, host)

    db, err := sql.Open("postgres", connURL)
    if err != nil {
        fmt.Println(err)
    }

    return db
}
