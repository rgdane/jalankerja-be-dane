
package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type ProjectSquadRepository interface {
	WithTx(tx *gorm.DB) ProjectSquadRepository
	WithPreloads(preloads ...string)ProjectSquadRepository

	InsertProjectSquad(input *models.ProjectSquad) (*models.ProjectSquad, error)
	UpdateProjectSquad(id int64, updates map[string]interface{}) (*models.ProjectSquad, error)
	RemoveProjectSquad(data *models.ProjectSquad) (*models.ProjectSquad, error)
	RemoveProjectSquadByID(id int64) error
	
	FindProjectSquad() ([]models.ProjectSquad, error)
	FindProjectSquadByID(id int64) (*models.ProjectSquad, error)
}
