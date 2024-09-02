package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"uln/src/handlers"
	ln "uln/src/services/ln"
)

const PORT int = 8080

type ulnApp struct {
    
}

func main() {
    app := echo.New()

    app.Use(middleware.Logger())

    homeHandler := handlers.HomeHandler{}
    app.GET("/", homeHandler.Handle)
    
    // admin := app.Group("/admin", )
    // admin.GET("/ln",)

    ln.Init()
    // ln := app.Group("ln", m ...echo.MiddlewareFunc)
    app.POST("/ln/create", ln.PostShortlink)
    app.POST("/ln/info", ln.PostShortlinkInfo)
    // app.GET("/ln/info/:short", ln.PostShortlinkInfo)
    app.DELETE("/ln/delete", ln.DeleteShortlink)
    app.GET("/:path", ln.GetRedirect) 

    app.Start(":8080")
}
