package routes

import (
	"first-rest-api-go/controller/auth"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// route auth
	router.POST("/api/register", controller_auth.Register)
	router.POST("/api/login", controller_auth.Login)

	return router
}