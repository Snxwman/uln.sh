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

func makeShortlink(rawURL string) (*shortlink, error) {
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
    
    path := makePath(7)
    shortURL, err := url.Parse(BASE_URL + "/" + path)
    if err != nil {
        return nil, CouldNotMakePathError{}
    }
    
    shortlink := shortlink {
        fullURL: *fullURL,
        shortURL: *shortURL,
        active: true,
        expiration: time.Now().AddDate(1, 0, 0),
        reserved: false,
        redirectReqs: 0,
        infoReqs: 0,
        options: shortURLCreationOptions {
            shortType: "random",
        },
        creationMetadata: models.CreationMetadata {
            CreatedAt: time.Now(),
            CreatedByUser: "",
            CreatedByIP: nil, 
            CreatedVia: "",
            InitialCreation: true,
        },
    }
    return &shortlink, nil
} 

