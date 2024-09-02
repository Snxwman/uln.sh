package models

import (
	"net"
	"time"

	"github.com/labstack/echo/v4"

	"uln/src/util"
)

type CreationMetadata struct {
    CreatedAt        time.Time
    CreatedByUser    string  // TODO: Change to User type when impl'd
    CreatedByIP      net.IP
    CreatedVia       string
    InitialCreation  bool
}

func MakeCreationMetadata(c echo.Context, initial bool) CreationMetadata {
    var via string
    if util.RequestViaCli(c) {
        via = "cli"
    } else {
        via = "web"
    }

    return CreationMetadata {
        CreatedAt: time.Now(),
        CreatedByUser: "anonymous",
        CreatedByIP: net.ParseIP(c.RealIP()),
        CreatedVia: via,
        InitialCreation: initial,
    }
}
