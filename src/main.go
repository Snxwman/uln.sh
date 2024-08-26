package main

import (
	"github.com/labstack/echo/v4"

	"uln/src/handlers"
	ln "uln/src/services/ln"
)

func main() {
    app := echo.New()

    homeHandler := handlers.HomeHandler{}
    app.GET("/", homeHandler.Handle)
    
    app.GET("/ln/create", ln.GetShortLink)
    app.GET("/ln/info", ln.GetShortLinkInfo)
    app.GET("/:short", ln.Redirect) 

    app.Start(":8080")
}
