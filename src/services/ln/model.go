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
    reserved         bool
    expiration       time.Time
    redirectReqs     int
    infoReqs         int
    lastAccessed     time.Time
    options          shortlinkCreationOptions
    creationMetadata models.CreationMetadata 
    // managementToken  string  // For anonymous management
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
    
    path := makePath(7)
    shortURL, err := url.Parse(BASE_URL + "/" + path)
    if err != nil {
        return nil, CouldNotMakePathError{}
    }
    
    shortlink := shortlink {
        fullURL: *fullURL,
        shortURL: *shortURL,
        active: true,
        reserved: false,
        expiration: time.Now().AddDate(1, 0, 0),
        redirectReqs: 0,
        infoReqs: 0,
        lastAccessed: time.Time{},
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
