package response

import (
	"first-rest-api-go/model"
)

type LoginResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}