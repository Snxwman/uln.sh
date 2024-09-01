package ln

import (
	"net/url"
	"time"
	"uln/src/models"
)

type Shortlink struct {
    fullURL          url.URL
    shortURL         url.URL
    active           bool
    expiration       time.Time
    reserved         bool
    redirectReqs     int
    infoReqs         int
    options          ShortURLCreationOptions
    creationMetadata models.CreationMetadata 
}

type ShortURLCreationOptions struct {
    shortType string  // TODO: Change to enum
}

