
package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type squadRepository struct {
	db *gorm.DB
}

func NewSquadRepository() sql.SquadRepository {
	return &squadRepository{db: config.DB}
}

func (repo *squadRepository) WithTx(tx *gorm.DB) sql.SquadRepository {
	return &squadRepository{db: tx}
}

func (repo *squadRepository) InsertSquad(squad *models.Squad) (*models.Squad, error) {
	if err := repo.db.Create(squad).Error; err != nil {
		return nil, err
	}
	return squad, nil
}

func (repo *squadRepository) UpdateSquads(squad *models.Squad) (*models.Squad, error) {
	if err := repo.db.Save(squad).Error; err != nil {
		return nil, err
	}
	return squad, nil
}

func (repo *squadRepository) RemoveSquad(squad *models.Squad) (*models.Squad, error) {
	if err := repo.db.Delete(squad).Error; err != nil {
		return nil, err
	}
	return squad, nil
}

func (repo *squadRepository) RemoveSquadByID(id int64) error {
	return repo.db.Delete(&models.Squad{}, id).Error
}

func (repo *squadRepository) FindSquad() ([]models.Squad, error) {
	var items []models.Squad
	if err := repo.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (repo *squadRepository) FindSquadByID(id int64) (*models.Squad, error) {
	var item models.Squad
	if err := repo.db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
		