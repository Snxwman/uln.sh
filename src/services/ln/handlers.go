package ln

import (
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
    rawURL := c.Request().FormValue("url")

    // Check if a shortlink already exists
    if path, exists := shortlinkExists(rawURL); exists {
        shortlink := lnApp.urls[path]

        if util.RequestViaCli(c) {
            return c.String(http.StatusOK, shortlink.shortURL.String() + "\n") 
        } else {
            c.Response().WriteHeader(http.StatusOK)
            return util.Render(c, ShortlinkTemplate(shortlink.shortURL.String())) 
        }
    }

    creationMetadata := models.MakeCreationMetadata(c, true)
    shortlink, err := makeShortlink(rawURL, creationMetadata)

    switch err.(type) {
    case EmptyURLError:
        return c.String(http.StatusBadRequest, err.Error())
    case CouldNotParseURLError:
        return c.String(http.StatusBadRequest, err.Error())
    case CouldNotMakePathError:
        return c.String(http.StatusBadRequest, err.Error())
    }

    err = registerShortlink(shortlink)
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    }

    if util.RequestViaCli(c) {
        return c.String(http.StatusCreated, shortlink.shortURL.String() + "\n") 
    } else {
        c.Response().WriteHeader(http.StatusCreated)
        return util.Render(c, ShortlinkTemplate(shortlink.shortURL.String())) 
    }
}

func PostShortlinkInfo(c echo.Context) error {
    url, err := url.Parse(c.Request().FormValue("url"))
    if err != nil {
        return c.String(http.StatusBadRequest, err.Error())
    } 

    path := strings.Trim(url.Path, "/")

    shortlink, ok := lnApp.urls[path]
    if !ok {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "No shortlink found"})
    }

    shortlink.infoReqs++
    lnApp.urls[path] = shortlink

    if util.RequestViaCli(c) {
        return c.String(http.StatusOK, shortlink.fullURL.String()) 
    } else {
        c.Response().WriteHeader(http.StatusOK)
        return util.Render(c, ShortlinkInfoTemplate(shortlink))
    }
}

func GetRedirect(c echo.Context) error {
    path := c.Param("path")
    shortlink, ok := lnApp.urls[path]
    if !ok {
        return c.String(http.StatusNotFound, "No shortlink found")
    }

    shortlink.redirectReqs++
    shortlink.lastAccessed = time.Now()
    lnApp.urls[path] = shortlink

    // Headers from google url shortener
    // https://stackoverflow.com/questions/47770376/why-does-url-shortening-service-send-response-with-http-status-codes-301-and-cac
    // https://stackoverflow.com/questions/9012456/using-301-303-307-redirects-for-dynamic-short-urls
    c.Response().Header().Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
    c.Response().Header().Add("Pragma", "no-cache")
    c.Response().Header().Add("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
    return c.Redirect(http.StatusPermanentRedirect, shortlink.fullURL.String())
}

func DeleteShortlink(c echo.Context) error {
    return nil
}
