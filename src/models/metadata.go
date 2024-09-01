package models

import (
    "net"
    "time"
)

type CreationMetadata struct {
    CreatedAt        time.Time
    CreatedByUser    string  // TODO: Change to User type when impl'd
    CreatedByIP      net.IP
    CreatedVia       string
    InitialCreation  bool
}
