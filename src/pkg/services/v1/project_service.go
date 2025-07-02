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

type ProjectService interface {
	WithTx(tx *gorm.DB) ProjectService

	CreateProject(input *models.Project) (*models.Project, error)
	UpdateProject(id int64, updates map[string]interface{}) (*models.Project, error)
	DeleteProject(id int64) error
	GetAllProjects() ([]models.Project, error)
	GetProjectByID(id int64) (*models.Project, error)
	GetDB() *gorm.DB
}

type projectService struct {
	repo sql.ProjectRepository
	tx   *gorm.DB
}

func NewProjectService(repo sql.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) WithTx(tx *gorm.DB) ProjectService {
	return &projectService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *projectService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *projectService) CreateProject(input *models.Project) (*models.Project, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertProject(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *projectService) UpdateProject(id int64, updates map[string]interface{}) (*models.Project, error) {
	if _, err := s.repo.FindProjectByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateProject(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *projectService) DeleteProject(id int64) error {
	err := s.repo.RemoveProjectByID(id)
	return gorm_err.TranslateGormError(err)
}

func (s *projectService) GetAllProjects() ([]models.Project, error) {
	data, err := s.repo.
		WithPreloads("ProjectSquad.HasSquad").FindProject()
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *projectService) GetProjectByID(id int64) (*models.Project, error) {
	data, err := s.repo.
		WithPreloads("ProjectSquad.HasSquad").FindProjectByID(id)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}
