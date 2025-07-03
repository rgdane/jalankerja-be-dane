
package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	WithTx(tx *gorm.DB) CategoryRepository
	WithPreloads(preloads ...string)CategoryRepository

	InsertCategory(input *models.Category) (*models.Category, error)
	UpdateCategory(id int64, updates map[string]interface{}) (*models.Category, error)
	RemoveCategory(data *models.Category) (*models.Category, error)
	RemoveCategoryByID(id int64) error
	
	FindCategory() ([]models.Category, error)
	FindCategoryByID(id int64) (*models.Category, error)
}
