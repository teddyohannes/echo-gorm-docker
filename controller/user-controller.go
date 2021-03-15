package controller

import (
	"echo-gorm-docker/repository"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	Update(context echo.Context)
	Profile(context echo.Context)
}

func Update(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Update!")
}

func Profile(c echo.Context) error {
	// return c.String(http.StatusOK, "Hello, Profile!")
	id := c.FormValue("id")
	fmt.Println("controller profile")
	result, err := repository.ProfileUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
