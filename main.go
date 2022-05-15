package main

import (
	"eu-and-vk-analysis/backend"
)

// @title User API documentation
// @version 1.0.0
// @host localhost:8000

func main() {
	app := app.NewApp()
	app.Run()
}