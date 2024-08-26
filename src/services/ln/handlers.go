package ln

import (
    "fmt"
    "github.com/labstack/echo/v4"
)

func GetShortLink(c echo.Context) error {
    fmt.Println(c)
    return nil
}

func GetShortLinkInfo(c echo.Context) error {
    return nil
}

func Redirect(c echo.Context) error {
    fmt.Println(c.Param("short"))
    return nil
}
