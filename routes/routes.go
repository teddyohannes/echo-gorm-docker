package routes

import (
	"echo-gorm-docker/controller"
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

	userRoutes := e.Group("api/user")
	{
		userRoutes.GET("/profile", controller.Profile)
		userRoutes.POST("/update", controller.Update)
	}
	return e
}
