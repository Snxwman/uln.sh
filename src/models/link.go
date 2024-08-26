package types

import (
    "net/url"
	"time"
)

type Link struct {
    original url.URL
    shortened url.URL
    createdBy User
    createdAt time.Time
}

func NewLink(raw string) (Link, error) {
    original, err := url.Parse(raw)
    if err != nil {
        return Link{}, err
    }

    // TODO:
    shortened, err := url.Parse("uln.sh/ot8EYbd")
    if err != nil {
        return Link{}, err
    }
    
    return Link {
        original: *original,
        shortened: *shortened,
        createdAt: time.Now(),
    }, nil
}

