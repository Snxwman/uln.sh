package models

import "time"

type User struct {
    id string
    ipAddresses []string
    username string
    email string
    banned bool
    banExpiration time.Time
    permissions permissions
}

type permissions struct {}
