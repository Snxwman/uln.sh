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
    
    ln.Init()
    app.POST("/ln/create", ln.PostShortLink)
    app.POST("/ln/info", ln.PostShortLinkInfo)
    app.GET("/:short", ln.GetRedirect) 

    app.Start(":8080")
}
