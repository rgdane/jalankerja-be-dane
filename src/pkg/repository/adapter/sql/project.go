
package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	WithTx(tx *gorm.DB) ProjectRepository
	WithPreloads(preloads ...string)ProjectRepository

	InsertProject(input *models.Project) (*models.Project, error)
	UpdateProject(id int64, updates map[string]interface{}) (*models.Project, error)
	RemoveProject(data *models.Project) (*models.Project, error)
	RemoveProjectByID(id int64) error
	
	FindProject() ([]models.Project, error)
	FindProjectByID(id int64) (*models.Project, error)
}
