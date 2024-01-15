package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vincentmegia/go-login-htmx/handlers"
)

func main() {
	app := echo.New()
	app.GET("/", handlers.Handle)
	app.Logger.Fatal(app.Start(":8080"))
}
