package main

import (
	"first-rest-api-go/config"
	"github.com/gin-gonic/gin"
)

func main() {

	//load file .env
	config.LoadEnv()

	//init database
	config.InitDB()

	//init gin
	router := gin.Default()

	// create route GET /
	router.GET("/", func(c *gin.Context) {

		//return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// run server on port .env APP_PORT or 3000
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}