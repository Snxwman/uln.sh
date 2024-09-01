package util

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
    return component.Render(c.Request().Context(), c.Response())
}

func RequestViaCli(c echo.Context) bool {
    userAgent := c.Request().Header.Get("User-Agent")
    if strings.Contains(userAgent, "curl") || strings.Contains(userAgent, "HTTPie") {
        return true
    } else {
        return false
    }
}
