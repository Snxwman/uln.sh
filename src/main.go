package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"uln/src/handlers"
	ln "uln/src/services/ln"
)

const PORT int = 8080

func main() {

    app := echo.New()

    app.Use(middleware.Logger())

    homeHandler := handlers.HomeHandler{}
    app.GET("/", homeHandler.Handle)
    
    // admin := app.Group("/admin", )

    ln.Init()
    // ln := app.Group("ln", m ...echo.MiddlewareFunc)
    app.POST("/ln/create", ln.PostShortLink)
    app.POST("/ln/info", ln.PostShortLinkInfo)
    app.GET("/:short", ln.GetRedirect) 

    app.Start(":8080")
}
