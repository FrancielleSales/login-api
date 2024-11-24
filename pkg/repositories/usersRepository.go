package repositories

import (
	"login-api/pkg/models"
	"gorm.io/gorm"
	"errors"
)

var (
	ErrNotFound = errors.New("record not found")
)

type UserRepository interface {
	CreateUser(user *models.Users) error
	GetUserByEmail(email string) (*models.Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.Users) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
