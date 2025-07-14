package ln

import (
	"net/url"
	"time"

	"uln/src/models"
)

type shortlink struct {
    FullURL          url.URL                    `db:"full_url"`
    ShortURL         url.URL                    `db:"short_url"`
    Active           bool                       `db:"active"`
    Reserved         bool                       `db:"reserved"`
    Expiration       time.Time                  `db:"expiration"`
    RedirectReqs     int                        `db:"redirect_reqs"`
    InfoReqs         int                        `db:"info_reqs"`
    LastAccessed     time.Time                  `db:"last_accessed"`
    options          shortlinkCreationOptions
    creationMetadata models.CreationMetadata 
    // managementToken  string  // For anonymous management
    // cacheInvalidate bool
    // tableId uuid
}

type shortlinkCreationOptions struct {
    shortlinkType string  // TODO: Change to enum
    unique bool
}

type EmptyURLError struct {} 
func (e EmptyURLError) Error() string {
    return "Provided URL was empty"
}

type CouldNotParseURLError struct {}
func (e CouldNotParseURLError) Error() string {
    return "Could not parse provided URL"
}

type CouldNotMakePathError struct {}
func (e CouldNotMakePathError) Error() string {
    return "Could not make path"
}

func makeShortlink(rawURL string, creationMetadata models.CreationMetadata) (*shortlink, error) {
    if rawURL == "" {
        return nil, EmptyURLError{}
    }
    
    fullURL, err := url.Parse(rawURL)
    if err != nil {
        return nil, CouldNotParseURLError{}
    }

    // TODO: 
    //     - Check URL is valid
    //     - Make sure (minimally) scheme, domain, and TLD are populated
    
    path := makeRandomPath(7)
    shortURL, err := url.Parse(BASE_URL + "/" + path)
    if err != nil {
        return nil, CouldNotMakePathError{}
    }
    
    shortlink := shortlink {
        FullURL: *fullURL,
        ShortURL: *shortURL,
        Active: true,
        Reserved: false,
        Expiration: time.Now().AddDate(1, 0, 0),
        RedirectReqs: 0,
        InfoReqs: 0,
        LastAccessed: time.Time{},
        options: makeShortlinkCreationOptions(), 
        creationMetadata: creationMetadata,
    }
    return &shortlink, nil
} 

func makeShortlinkCreationOptions() shortlinkCreationOptions {
    return shortlinkCreationOptions {
        shortlinkType: "random",
        unique: false,
    }
}
