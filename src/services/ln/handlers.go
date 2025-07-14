package ln

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"uln/src/models"
	"uln/src/util"
)

// TODO: Create unique override
func PostShortlink(c echo.Context) error {
    // Get the URL to shorten from the request form
    rawURL := c.Request().FormValue("url")

    shortlinks, exists := tryGetShortlinkForUrl(rawURL)
    if exists { 
        shortlink := shortlinks[0]
        fmt.Println(shortlink)
        if util.RequestViaCli(c) {
            return c.String(http.StatusOK, shortlink.ShortURL.String() + "\n") 
        } else {
            c.Response().WriteHeader(http.StatusOK)
            return util.Render(c, ShortlinkTemplate(shortlink.ShortURL.String())) 
        }
    }

    isInitialCreation := true
    creationMetadata := models.MakeCreationMetadata(c, isInitialCreation)
    // shortlinkCreationOptions := makeShortlinkCreationOptions()
    shortlink, err := makeShortlink(rawURL, creationMetadata)

    // Handle errors while making the shortlink
    switch err.(type) {
    case EmptyURLError:
        return c.String(http.StatusBadRequest, err.Error())
    case CouldNotParseURLError:
        return c.String(http.StatusBadRequest, err.Error())
    case CouldNotMakePathError:
        return c.String(http.StatusBadRequest, err.Error())
    }

    // Add shortlink to cache
    err = registerShortlink(shortlink)
    if err != nil {
        log.Printf("Error registering shortlink: %s", err.Error())
        return c.String(http.StatusBadRequest, err.Error())
    }

    // Commit shortlink to database
    err = lnApp.db.insertShortlink(shortlink)
    if err != nil {
        log.Printf("Error inserting shortlink: %s", err.Error())
        return errors.New("Error while inserting shortlink")
    }

    // Return the result to user
    if util.RequestViaCli(c) {
        return c.String(http.StatusCreated, shortlink.ShortURL.String() + "\n") 
    } else {
        c.Response().WriteHeader(http.StatusCreated)
        return util.Render(c, ShortlinkTemplate(shortlink.ShortURL.String())) 
    }
}

func PostShortlinkInfo(c echo.Context) error {
    url, err := url.Parse(c.Request().FormValue("url"))
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    } 

    path := strings.Trim(url.Path, "/")

    shortlink, inCache := lnApp.cache.contains(path)
    if !inCache {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "No shortlink found"})
    }

    shortlink.InfoReqs++
    lnApp.cache.insert(shortlink)

    if util.RequestViaCli(c) {
        return c.String(http.StatusOK, shortlink.FullURL.String()) 
    } else {
        c.Response().WriteHeader(http.StatusOK)
        return util.Render(c, ShortlinkInfoTemplate(shortlink))
    }
}

func GetRedirect(c echo.Context) error {
    path := c.Param("path")
    shortlink, exists := tryGetShortlinkByPath(path)
    if exists {
        return c.String(http.StatusNotFound, "No shortlink found")
    }

    shortlink.RedirectReqs++
    shortlink.LastAccessed = time.Now()
    lnApp.cache.insert(shortlink)

    // Headers from google url shortener
    // https://stackoverflow.com/questions/47770376/why-does-url-shortening-service-send-response-with-http-status-codes-301-and-cac
    // https://stackoverflow.com/questions/9012456/using-301-303-307-redirects-for-dynamic-short-urls
    c.Response().Header().Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
    c.Response().Header().Add("Pragma", "no-cache")
    c.Response().Header().Add("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
    return c.Redirect(http.StatusPermanentRedirect, shortlink.FullURL.String())
}

func DeleteShortlink(c echo.Context) error {
    return nil
}
