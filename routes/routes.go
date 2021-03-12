package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/go", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go!")
	})
	return e
}
