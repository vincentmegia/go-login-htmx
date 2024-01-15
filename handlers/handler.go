package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincentmegia/go-login-htmx/repositories"
)

func Handle(c echo.Context) error {
	users, error := repositories.GetAllUsers()
	if error != nil {
		fmt.Println("An error occured whilte fetching all users", error)
		return error
	}

	return c.JSONPretty(http.StatusOK, users, "  ")
}
