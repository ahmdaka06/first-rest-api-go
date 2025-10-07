package controllers

import (
	"net/http"
	"first-rest-api-go/database"
	"first-rest-api-go/helper"
	"first-rest-api-go/model"
	"first-rest-api-go/structs"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func Login(c *gin.Context) {

	// Inisialisasi struct untuk menampung data dari request
	var req = structs.UserLoginRequest{}
	var user = model.User{}

	// Validasi input dari request body menggunakan ShouldBindJSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	// Cari user berdasarkan username yang diberikan di database
	// Jika tidak ditemukan, kirimkan respons error Unauthorized
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User Not Found",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	// Bandingkan password yang dimasukkan dengan password yang sudah di-hash di database
	// Jika tidak cocok, kirimkan respons error Unauthorized
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}

	// Jika login berhasil, generate token untuk user
	token, err := helper.GenerateJWT(user.Id, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to generate token",
			Errors:  helper.TranslateErrorMessage(err),
		})
		return
	}
	

	// Kirimkan response sukses dengan status OK dan data user serta token
	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "Login Success",
		Data: structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			Token:     &token,
		},
	})
}