package types

type permissions struct {}

type User struct {
    id string
    ipAddresses []string
    username string
    email string
    permissions permissions
}
