package ln

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo/v4"

	"uln/src/models"
	"uln/src/util"
)

func PostShortLink(c echo.Context) error {
    fmt.Println(c.FormValue("url"))

    rawURL := c.Request().FormValue("url")
    if rawURL == "" {
        return c.String(http.StatusBadRequest, "No URL provided")    
    }
    
    fullURL, err := url.Parse(c.Request().FormValue("url"))
    if err != nil {
        return c.String(http.StatusBadRequest, "Invalid URL")
    }
    
    shortURL, err := url.Parse("http://uln.sh/" + makePath(7))
    if err != nil {
        log.Fatal("failed to parse short url")
    }
    
    via := ""
    if util.RequestViaCli(c) {
        via = "cli"
    } else {
        via = "web"
    }

    ln := shortlink {
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
            CreatedByIP: net.ParseIP(c.RealIP()),
            CreatedVia: via,
            InitialCreation: true,
        },
    }

    registerShortlink(ln)
    fmt.Println(ln)

    if util.RequestViaCli(c) {
        return c.String(http.StatusCreated, ln.shortURL.String() + "\n") 
    } else {
        c.Response().WriteHeader(http.StatusCreated)
        return util.Render(c, ShortLink(ln.shortURL.String())) 
    }
}


func PostShortLinkInfo(c echo.Context) error {
    fmt.Println(c.FormValue("url"))
    return c.String(http.StatusOK, "info") 
}

func Redirect(c echo.Context) error {
    fmt.Println(c.Param("short"))
    path := c.Param("short")
    fullURL := lns[path].fullURL
    return c.Redirect(http.StatusTemporaryRedirect, fullURL.String())
}
