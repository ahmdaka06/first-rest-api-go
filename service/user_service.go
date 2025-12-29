package service

import (
	"errors"
	"first-rest-api-go/helper"
	"first-rest-api-go/model"
	"first-rest-api-go/repository"
	"first-rest-api-go/structs"
)

type UserService interface {
	Login(email string, password string) (*structs.UserResponse, error)
	Register(name string, email string, password string) (*structs.UserResponse, error)
	Create(name string, email string, password string) (*structs.UserResponse, error)
	FindAll() ([]structs.UserResponse, error)
	FindById(id uint) (*structs.UserResponse, error)
	UpdateById(id uint, name string, email string, password string) (*structs.UserResponse, error)
	DeleteById(id uint) (*structs.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) Login(email string, password string) (*structs.UserResponse, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("User not found")
	}
	if !helper.CheckPassword(user.Password, password) {
		return nil, errors.New("Email or password is wrong")
	}

	token, err := helper.GenerateJWT(user.Id, user.Email)
	if err != nil {
		return nil, errors.New("Login failed")
	}

	return &structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
		Token:     &token,
	}, nil
}

func (s *userService) Register(name string, email string, password string) (*structs.UserResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil && existingUser.Id != 0 {
		return nil, errors.New("Email already registered")
	}

	// Hash password
	hashedPassword := helper.HashPassword(password)

	// Create user model
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	// Save to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("Failed to create user")
	}

	return &structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (s *userService) Create(name string, email string, password string) (*structs.UserResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(email)
	if existingUser != nil && existingUser.Id != 0 {
		return nil, errors.New("Email already registered")
	}

	// Hash password
	hashedPassword := helper.HashPassword(password)

	// Create user model
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	// Save to database
	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("Failed to create user")
	}

	return &structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (s *userService) FindAll() ([]structs.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, errors.New("Failed to get users")
	}

	var userResponses []structs.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, structs.UserResponse{
			Id:        user.Id,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}
	return userResponses, nil
}

func (s *userService) FindById(id uint) (*structs.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return &structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (s *userService) UpdateById(id uint, name string, email string, password string) (*structs.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	user.Name = name
	user.Email = email
	user.Password = helper.HashPassword(password)
	if err := s.userRepo.UpdateById(id, user); err != nil {
		return nil, err
	}
	return &structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (s *userService) DeleteById(id uint) (*structs.UserResponse, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	if err := s.userRepo.DeleteById(id); err != nil {
		return nil, err
	}
	return &structs.UserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}
