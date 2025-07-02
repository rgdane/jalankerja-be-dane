package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type squadRepository struct {
	db       *gorm.DB
	preloads []string
}

func NewSquadRepository() sql.SquadRepository {
	return &squadRepository{db: config.DB}
}

func (repo *squadRepository) WithTx(tx *gorm.DB) sql.SquadRepository {
	return &squadRepository{
		db:       tx,
		preloads: repo.preloads,
	}
}

func (repo *squadRepository) WithPreloads(preloads ...string) sql.SquadRepository {
	return &squadRepository{
		db:       repo.db,
		preloads: preloads,
	}
}

func (repo *squadRepository) applyPreloads(db *gorm.DB) *gorm.DB {
	for _, relation := range repo.preloads {
		db = db.Preload(relation)
	}
	return db
}

func (repo *squadRepository) InsertSquad(data *models.Squad) (*models.Squad, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *squadRepository) UpdateSquad(id int64, updates map[string]interface{}) (*models.Squad, error) {
	if err := repo.db.Model(&models.Squad{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	var updated models.Squad
	if err := repo.db.First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (repo *squadRepository) RemoveSquad(data *models.Squad) (*models.Squad, error) {
	if err := repo.db.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *squadRepository) RemoveSquadByID(id int64) error {
	return repo.db.Delete(&models.Squad{}, id).Error
}

func (repo *squadRepository) FindSquad() ([]models.Squad, error) {
	var items []models.Squad
	db := repo.applyPreloads(repo.db.Model(&models.Squad{}))

	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *squadRepository) FindSquadByID(id int64) (*models.Squad, error) {
	var item models.Squad
	db := repo.applyPreloads(repo.db)

	if err := db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
