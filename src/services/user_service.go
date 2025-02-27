package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/mutinho/src/dtos"
	"github.com/mutinho/src/models"
	"github.com/mutinho/src/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	id := uuid.New()

	user := &models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Telefono: req.Telefono,
		Rol:      models.Role(req.Rol),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Telefono:  user.Telefono,
		Rol:       string(user.Rol),
		Slug:      user.Slug,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil

}

func (s *UserService) GetUserByEmail(email string) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Telefono:  user.Telefono,
		Rol:       string(user.Rol),
		Slug:      user.Slug,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) GetUserBySlug(slug string) (*dtos.UserResponse, error) {
	user, err := s.userRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}

	return &dtos.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		Telefono:  user.Telefono,
		Rol:       string(user.Rol),
		Slug:      user.Slug,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) GetAllUsers() ([]dtos.UserResponse, error) {
	users, err := s.userRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var userResponses []dtos.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, dtos.UserResponse{
			Name:      user.Name,
			Email:     user.Email,
			Telefono:  user.Telefono,
			Rol:       string(user.Rol),
			Slug:      user.Slug,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		})
	}

	return userResponses, nil
}
