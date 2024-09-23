package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"

	"uln/src/handlers"
	ln "uln/src/services/ln"
	"uln/src/store"
)

const PORT int = 8080

type ulnApp struct {
    // config config
    db *sql.DB
}

func main() {
    uln := ulnApp {
        // config: config {},
        db: store.Init(),
    }
    defer uln.db.Close()

    app := echo.New()
    // app.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")

    app.Use(middleware.Recover())
    app.Use(middleware.Logger())

    // app.Pre(middleware.HTTPSRedirect())
    // app.Pre(middleware.HTTPSNonWWWRedirect())

    homeHandler := handlers.HomeHandler{}
    app.GET("/", homeHandler.Handle)
    
    // admin := app.Group("/admin", )
    // admin.GET("/ln",)

    ln.Init(uln.db)
    // ln := app.Group("ln", m ...echo.MiddlewareFunc)
    app.POST("/ln/create", ln.PostShortlink)
    app.POST("/ln/info", ln.PostShortlinkInfo)
    app.GET("/:path", ln.GetRedirect) 
    // app.GET("/:path/info", ln.PostShortlinkInfo)
    app.DELETE("/:path/delete", ln.DeleteShortlink)

    s := &http2.Server {
        MaxConcurrentStreams: 250,
        MaxReadFrameSize:     1048576,
        IdleTimeout:          10 * time.Second,
    }

    err := app.StartH2CServer(fmt.Sprintf(":%d", PORT), s)
    if err != nil {
        fmt.Println(err)
    }
}
