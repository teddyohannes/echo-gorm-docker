package main

import (
	"echo-gorm-docker/config"
	"echo-gorm-docker/routes"
)

func main() {
	config.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
