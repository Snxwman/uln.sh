package handlers

import (
	"github.com/labstack/echo/v4"

	"uln/src/templates"
)

type HomeHandler struct {}

func (h HomeHandler) Handle(c echo.Context) error {
    return render(c, templates.Home())
}

func GetHome(c echo.Context) error {
    return nil
}
