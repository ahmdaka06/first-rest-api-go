package main

import (
	"first-rest-api-go/config"
	"first-rest-api-go/routes"
)

func main() {

	//load file .env
	config.LoadEnv()

	//init database
	config.InitDB()

	//setup router
	r := routes.SetupRouter()

	// run server on port .env APP_PORT or 3000
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}