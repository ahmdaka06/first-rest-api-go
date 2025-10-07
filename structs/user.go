package structs

type UserResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omitempty"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique;not null"`
	Password string `json:"password,omitempty"`
}

type UserLoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}