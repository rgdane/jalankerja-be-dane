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

type ProjectSquadService interface {
	WithTx(tx *gorm.DB) ProjectSquadService

	CreateProjectSquad(input *models.ProjectSquad) (*models.ProjectSquad, error)
	UpdateProjectSquad(id int64, updates map[string]interface{}) (*models.ProjectSquad, error)
	DeleteProjectSquad(id int64) error
	GetAllProjectSquads() ([]models.ProjectSquad, error)
	GetProjectSquadByID(id int64) (*models.ProjectSquad, error)
	GetDB() *gorm.DB
}

type projectSquadService struct {
	repo sql.ProjectSquadRepository
	tx   *gorm.DB
}

func NewProjectSquadService(repo sql.ProjectSquadRepository) ProjectSquadService {
	return &projectSquadService{repo: repo}
}

func (s *projectSquadService) WithTx(tx *gorm.DB) ProjectSquadService {
	return &projectSquadService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *projectSquadService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *projectSquadService) CreateProjectSquad(input *models.ProjectSquad) (*models.ProjectSquad, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertProjectSquad(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *projectSquadService) UpdateProjectSquad(id int64, updates map[string]interface{}) (*models.ProjectSquad, error) {
	if _, err := s.repo.FindProjectSquadByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateProjectSquad(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *projectSquadService) DeleteProjectSquad(id int64) error {
	err := s.repo.RemoveProjectSquadByID(id)
	return gorm_err.TranslateGormError(err)
}

func (s *projectSquadService) GetAllProjectSquads() ([]models.ProjectSquad, error) {
	data, err := s.repo.
		WithPreloads().FindProjectSquad()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *projectSquadService) GetProjectSquadByID(id int64) (*models.ProjectSquad, error) {
	data, err := s.repo.
		WithPreloads().FindProjectSquadByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}
