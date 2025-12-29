package mapper

import (
	"first-rest-api-go/model"
	"first-rest-api-go/structs"
)

func ToUserResponse(user []model.User) []structs.UserResponse {
	var userResponses []structs.UserResponse
	for _, user := range user {
		userResponses = append(userResponses, structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}
	return userResponses
}