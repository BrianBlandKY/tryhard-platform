package main

import (
	"fmt"
	"log"
	"mime"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	// _ "github.com/labstack/echo/engine"
	// std "github.com/labstack/echo/engine/standard"
	cors "github.com/rs/cors"

	_ "net"
	_ "time"

	config "tryhard-platform/config"
	routes "tryhard-platform/routes"
	socket "tryhard-platform/server"
)

// Secure Route Handlers? API_KEY? Token? Cookie?
func init() {
	/* Add support for Markdown mime-types */
	mime.AddExtensionType(".markdown", "text/markdown")
	mime.AddExtensionType(".md", "text/markdown")
}

func StartServer(cfg config.Config) {
	defer func() {
		crash := recover()
		if crash != nil {
			log.Printf("Application Crash. %s", crash)
		}
	}()

	/* Parse server host:port */
	api := cfg.Api
	addr := fmt.Sprintf("%v:%v", api.Address, api.Port)

	e := echo.New()

	e.Debug = true

	e.Use(mw.Logger())
	e.Use(mw.Gzip())
	e.Use(mw.Recover())

	// CORS
	e.Use(echo.WrapMiddleware(cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://*:3000",
			"http://*:3001",
			"http://*:8080",
			"http://*:8181",
			"http://*:8282"},
	}).Handler))

	socketServer := socket.NewServer("http://localhost/")
	go socketServer.Listen(e)

	// Custom Client Id
	//e.Use(ClientCookie())
	routes.Router(e, cfg)

	log.Printf("API Server Started %v", addr)
	e.Start(addr)
}
