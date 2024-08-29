package ln

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PostShortLink(c echo.Context) error {
    fmt.Println(c.FormValue("url"))
    shortLink := "https://uln.sh/Ai3Vjg0"
    time.Sleep(3 * time.Second)
    return c.String(http.StatusOK, shortLink)
}


func PostShortLinkInfo(c echo.Context) error {
    fmt.Println(c.FormValue("url"))
    return c.String(http.StatusOK, "info") 
}

func Redirect(c echo.Context) error {
    fmt.Println(c.Param("short"))
    return c.Redirect(http.StatusTemporaryRedirect, "https://google.com")
}
