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

type SquadService interface {
	WithTx(tx *gorm.DB) SquadService

	CreateSquad(input *models.Squad) (*models.Squad, error)
	UpdateSquad(id int64, updates map[string]interface{}) (*models.Squad, error)
	DeleteSquad(id int64) error
	GetAllSquads() ([]models.Squad, error)
	GetSquadByID(id int64) (*models.Squad, error)
	GetDB() *gorm.DB
}

type squadService struct {
	repo sql.SquadRepository
	tx   *gorm.DB
}

func NewSquadService(repo sql.SquadRepository) SquadService {
	return &squadService{repo: repo}
}

func (s *squadService) WithTx(tx *gorm.DB) SquadService {
	return &squadService{
		repo: s.repo.WithTx(tx),
		tx:   tx,
	}
}

func (s *squadService) GetDB() *gorm.DB {
	if s.tx != nil {
		return s.tx
	}
	return config.DB
}

func (s *squadService) CreateSquad(input *models.Squad) (*models.Squad, error) {
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	data, err := s.repo.InsertSquad(input)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *squadService) UpdateSquad(id int64, updates map[string]interface{}) (*models.Squad, error) {
	if _, err := s.repo.FindSquadByID(id); err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}

	updates["updated_at"] = time.Now()
	fmt.Println(updates)

	data, err := s.repo.UpdateSquad(id, updates)
	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *squadService) DeleteSquad(id int64) error {
	err := s.repo.RemoveSquadByID(id)
	return gorm_err.TranslateGormError(err)
}

func (s *squadService) GetAllSquads() ([]models.Squad, error) {
	data, err := s.repo.
		WithPreloads("HasCaptain", "ProjectSquad.HasProject").FindSquad()

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}

func (s *squadService) GetSquadByID(id int64) (*models.Squad, error) {
	data, err := s.repo.
		WithPreloads("HasCaptain", "ProjectSquad.HasProject").FindSquadByID(id)

	if err != nil {
		return nil, gorm_err.TranslateGormError(err)
	}
	return data, nil
}
