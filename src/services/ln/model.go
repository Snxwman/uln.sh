package ln

import (
	"net/url"
	"time"
	"uln/src/models"
)

type shortlink struct {
    fullURL          url.URL
    shortURL         url.URL
    active           bool
    expiration       time.Time
    reserved         bool
    redirectReqs     int
    infoReqs         int
    options          shortURLCreationOptions
    creationMetadata models.CreationMetadata 
}

type shortURLCreationOptions struct {
    shortType string  // TODO: Change to enum
}

