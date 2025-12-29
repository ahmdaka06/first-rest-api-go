package controller_auth

import (
	"net/http"
	"first-rest-api-go/helper"
	"first-rest-api-go/structs"
	"github.com/gin-gonic/gin"
	"first-rest-api-go/service"
)

type LoginController struct {
	userService service.UserService
}

func NewLoginController(userService service.UserService) LoginController {
	return LoginController{userService: userService}
}

func (c *LoginController) Login(ctx *gin.Context) {

	// init struct for request body
	var req = structs.UserLoginRequest{}

	// Validate request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	result, err := c.userService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: err.Error(),
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	// Send success response
	ctx.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: result,
	})
}