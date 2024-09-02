package ln

import (
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo/v4"

	"uln/src/util"
)

func PostShortLink(c echo.Context) error {
    shortlink, err := makeShortlink(c.Request().FormValue("url"))

    switch err.(type) {
    case EmptyURLError:
        return c.String(http.StatusBadRequest, err.Error())
    case CouldNotParseURLError:
        return c.String(http.StatusBadRequest, err.Error())
    case CouldNotMakePathError:
        return c.String(http.StatusBadRequest, err.Error())
    }

    var via string
    if util.RequestViaCli(c) {
        via = "cli"
    } else {
        via = "web"
    }
        
    shortlink.creationMetadata.CreatedVia = via
    shortlink.creationMetadata.CreatedByIP = net.ParseIP(c.RealIP())

    err = registerShortlink(*shortlink)
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


func PostShortLinkInfo(c echo.Context) error {
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
    path := c.Param("short")
    shortlink, ok := lnApp.urls[path]
    if !ok {
        return c.String(http.StatusNotFound, "No shortlink found")
    }

    shortlink.redirectReqs++
    lnApp.urls[path] = shortlink

    c.Response().Header().Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate")
    c.Response().Header().Add("Pragma", "no-cache")
    c.Response().Header().Add("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
    return c.Redirect(http.StatusPermanentRedirect, shortlink.fullURL.String())
}
