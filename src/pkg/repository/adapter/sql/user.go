
package sql

import (
	"jk-api/internal/database/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	WithTx(tx *gorm.DB) UserRepository

	InsertUser(input *models.User) (*models.User, error)
	UpdateUsers(input *models.User) (*models.User, error)
	RemoveUser(data *models.User) (*models.User, error)
	RemoveUserByID(id int64) error
	FindUser() ([]models.User, error)
	FindUserByID(id int64) (*models.User, error)
}
	