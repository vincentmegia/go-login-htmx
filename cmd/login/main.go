package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	config "github.com/vincentmegia/go-login-htmx/build"
	"github.com/vincentmegia/go-login-htmx/internal/models"
	"github.com/vincentmegia/go-login-htmx/web/components"
)

var Env1 = ""

func main() {
	//yyyy-dd-mm
	html := components.Html()
	server := echo.New()
	server.GET("/", func(e echo.Context) error {
		fmt.Println("request is being handled here")
		return html.Render(e.Request().Context(), e.Response().Writer)
	})
	server.POST("/api/v1/user", func(e echo.Context) error {
		var user = models.User{}
		error := e.Bind(&user)
		if error != nil {
			fmt.Println("error: ", error)
			return e.String(http.StatusBadRequest, "bad request")
		}
		fmt.Println("passed user data:\t", user)
		return nil
	})

	fmt.Println("Environment:\t", config.Environment)
	fmt.Println("Version:\t", config.Version)
	fmt.Println("Main Env:\t", Env1)

	// move port to config file
	fmt.Println("Listening on :3000")
	server.Logger.Fatal(server.Start(":3000"))
}
