package services

import (
	"login-api/pkg/models"
	"login-api/pkg/repositories"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type UserService interface {
	CreateUser(name, email, password string) (*models.Users, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(name, email, password string) (*models.Users, error) {
	existingUser, err := s.repo.GetUserByEmail(email)
	if err != nil && err != repositories.ErrNotFound {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.Users{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
