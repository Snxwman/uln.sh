package ln

import "strings"

type lnCache struct {
    shortlinks map[string]*shortlink     // Maps the shortened path -> shortlink
    reverseShortlinks map[string]*shortlink
}

func newLnCache() lnCache {
    return lnCache {
        shortlinks: make(map[string]*shortlink),
        reverseShortlinks: make(map[string]*shortlink),
    }
}

func (c *lnCache) contains(path string) (*shortlink, bool) {
    shortlink, ok := c.shortlinks[path]
    if !ok {
        return nil, false
    }

    return shortlink, true
}

func (c *lnCache) containsReverse(url string) (*shortlink, bool) {
    shortlink, ok := c.reverseShortlinks[url]
    if !ok {
        return nil, false
    }

    return shortlink, true
}

func (c *lnCache) insert(s *shortlink) {
    path := strings.Trim(s.ShortURL.Path, "/")
    c.shortlinks[path] = s
    c.reverseShortlinks[s.FullURL.String()] = s
}

func (c *lnCache) remove(path string) {
    url := c.shortlinks[path].FullURL.String()
    delete(c.shortlinks, path) 
    delete(c.reverseShortlinks, url) 
}

func (c *lnCache) load() {

}


