package controller_auth

import (
	"first-rest-api-go/helper"
	"first-rest-api-go/service"
	"first-rest-api-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	userService service.UserService
}

func NewRegisterController(userService service.UserService) RegisterController {
	return RegisterController{userService: userService}
}

func (c *RegisterController) Register(ctx *gin.Context) {

	// init struct for request body
	var req = structs.UserCreateRequest{}

	// Validate request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	result, err := c.userService.Register(req.Name, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: err.Error(),
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	// Kirimkan response sukses dengan status Created dan data user
	ctx.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User Created Successfully",
		Data:    result,
	})
}
