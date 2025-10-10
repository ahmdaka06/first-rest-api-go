package controller_auth

import (
	"net/http"
	"first-rest-api-go/database"
	"first-rest-api-go/helper"
	"first-rest-api-go/model"
	"first-rest-api-go/structs"
	"github.com/gin-gonic/gin"
)

// Register menangani proses registrasi user baru
func Register(c *gin.Context) {
	// Inisialisasi struct untuk menangkap data request
	var req = structs.UserCreateRequest{}

	// Validasi request JSON menggunakan binding dari Gin
	if err := c.ShouldBindJSON(&req); err != nil {
		// Jika validasi gagal, kirimkan response error
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validasi Errors",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: helper.HashPassword(req.Password),
	}

	// Simpan data user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		if helper.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: false,
				Message: "Duplicate entry error",
				Errors:  helper.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: false,
				Message: "Failed to create user",
				Errors:  helper.TranslateErrorMessage(err),
			})
		}
		return
	}
	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}