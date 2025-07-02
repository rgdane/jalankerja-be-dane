package services

import (
	"errors"
	"fmt"
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"gorm.io/gorm"
)

type UserService interface {
	WithTx(tx *gorm.DB) UserService

	CreateUser(input *models.User) (*models.User, error)
	UpdateUser(id int64, input *models.User) (*models.User, error)
	DeleteUser(id int64) error
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int64) (*models.User, error)
	GetDB() *gorm.DB
}

type userService struct {
	repo sql.UserRepository
	tx   *gorm.DB
}

func NewUserService(repo sql.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) WithTx(tx *gorm.DB) UserService {
	return &userService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *userService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *userService) CreateUser(input *models.User) (*models.User, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertUser(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	return data, nil
}

func (s *userService) UpdateUser(id int64, input *models.User) (*models.User, error) {
	_, err := s.repo.FindUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	input.ID = (id)
	input.UpdatedAt = time.Now()
	fmt.Print(input)

	return s.repo.UpdateUsers(input)
}

func (s *userService) DeleteUser(id int64) error {
	return s.repo.RemoveUserByID(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.FindUser()
}

func (s *userService) GetUserByID(id int64) (*models.User, error) {
	return s.repo.FindUserByID(id)
}
