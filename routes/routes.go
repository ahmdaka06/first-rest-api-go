package routes

import (
	controller_auth "first-rest-api-go/controller/auth"
	"first-rest-api-go/database"
	"first-rest-api-go/repository"
	"first-rest-api-go/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Initialize repository
	userRepo := repository.NewUserRepository(database.DB)

	// Initialize service
	userService := service.NewUserService(userRepo)

	// Initialize controllers
	loginController := controller_auth.NewLoginController(userService)
	registerController := controller_auth.NewRegisterController(userService)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// route auth
	router.POST("/api/register", registerController.Register)
	router.POST("/api/login", loginController.Login)

	return router
}
