
package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() sql.UserRepository {
	return &userRepository{db: config.DB}
}

func (repo *userRepository) WithTx(tx *gorm.DB) sql.UserRepository {
	return &userRepository{db: tx}
}

func (repo *userRepository) InsertUser(data *models.User) (*models.User, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *userRepository) UpdateUsers(data *models.User) (*models.User, error) {
	if err := repo.db.Save(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *userRepository) RemoveUser(data *models.User) (*models.User, error) {
	if err := repo.db.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *userRepository) RemoveUserByID(id int64) error {
	return repo.db.Delete(&models.User{}, id).Error
}

func (repo *userRepository) FindUser() ([]models.User, error) {
	var items []models.User
	if err := repo.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *userRepository) FindUserByID(id int64) (*models.User, error) {
	var item models.User
	if err := repo.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
