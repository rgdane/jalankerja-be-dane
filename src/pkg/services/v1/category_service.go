package services

import (
	"fmt"
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/internal/errors/gorm_err"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"gorm.io/gorm"
)

type CategoryService interface {
	WithTx(tx *gorm.DB) CategoryService

	CreateCategory(input *models.Category) (*models.Category, error)
	UpdateCategory(id int64, updates map[string]interface{}) (*models.Category, error)
	DeleteCategory(id int64) error
	GetAllCategorys() ([]models.Category, error)
	GetCategoryByID(id int64) (*models.Category, error)
	GetDB() *gorm.DB
}

type categoryService struct {
	repo sql.CategoryRepository
	tx   *gorm.DB
}

func NewCategoryService(repo sql.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) WithTx(tx *gorm.DB) CategoryService {
	return &categoryService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *categoryService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *categoryService) CreateCategory(input *models.Category) (*models.Category, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertCategory(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *categoryService) UpdateCategory(id int64, updates map[string]interface{}) (*models.Category, error) {
	if _, err := s.repo.FindCategoryByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateCategory(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *categoryService) DeleteCategory(id int64) error {
	err := s.repo.RemoveCategoryByID(id)
	return gorm_err.TranslateGormError(err)
}

func (s *categoryService) GetAllCategorys() ([]models.Category, error) {
	data, err := s.repo.
		WithPreloads().FindCategory()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *categoryService) GetCategoryByID(id int64) (*models.Category, error) {
	data, err := s.repo.
		WithPreloads().FindCategoryByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}
