package main

import "echo-gorm-docker/routes"

// "github.com/labstack/echo"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
