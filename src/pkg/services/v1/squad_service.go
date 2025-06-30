package services

import (
	"errors"
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"
	"time"

	"gorm.io/gorm"
)

type SquadService interface {
	WithTx(tx *gorm.DB) SquadService

	CreateSquad(input *models.Squad) (*models.Squad, error)
	UpdateSquad(id int64, input *models.Squad) (*models.Squad, error)
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

	return s.repo.InsertSquad(input)
}

func (s *squadService) UpdateSquad(id int64, input *models.Squad) (*models.Squad, error) {
	existing, err := s.repo.FindSquadByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("squad not found")
		}
		return nil, err
	}

	if input.Name != "" {
		existing.Name = input.Name
	}
	existing.UpdatedAt = time.Now()

	return s.repo.UpdateSquads(existing)
}

func (s *squadService) DeleteSquad(id int64) error {
	return s.repo.RemoveSquadByID(id)
}

func (s *squadService) GetAllSquads() ([]models.Squad, error) {
	return s.repo.FindSquad()
}

func (s *squadService) GetSquadByID(id int64) (*models.Squad, error) {
	return s.repo.FindSquadByID(id)
}
