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
    
    app.POST("/ln/create", ln.PostShortLink)
    app.POST("/ln/info", ln.PostShortLinkInfo)
    app.GET("/:short", ln.Redirect) 

    app.Start(":8080")
}
