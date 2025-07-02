package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type SquadRepository interface {
	WithTx(tx *gorm.DB) SquadRepository
	WithPreloads(preloads ...string) SquadRepository

	InsertSquad(input *models.Squad) (*models.Squad, error)
	UpdateSquad(id int64, updates map[string]interface{}) (*models.Squad, error)
	RemoveSquad(data *models.Squad) (*models.Squad, error)
	RemoveSquadByID(id int64) error

	FindSquad() ([]models.Squad, error)
	FindSquadByID(id int64) (*models.Squad, error)
}
