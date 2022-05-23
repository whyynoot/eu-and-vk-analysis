package main

import (
	"eu-and-vk-analysis/backend"
)

// @title EU and VK Analytics API documentation
// @version 1.0.0
// @host euandvkanalysis.herokuapp.com
// @BasePath /

func main() {
	app := app.NewApp()
	app.Run()
}
