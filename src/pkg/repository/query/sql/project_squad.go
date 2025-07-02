package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type projectSquadRepository struct {
	db       *gorm.DB
	preloads []string
}

func NewProjectSquadRepository() sql.ProjectSquadRepository {
	return &projectSquadRepository{db: config.DB}
}

func (repo *projectSquadRepository) WithTx(tx *gorm.DB) sql.ProjectSquadRepository {
	return &projectSquadRepository{
		db:       tx,
		preloads: repo.preloads,
	}
}

func (repo *projectSquadRepository) WithPreloads(preloads ...string) sql.ProjectSquadRepository {
	return &projectSquadRepository{
		db:       repo.db,
		preloads: preloads,
	}
}

func (repo *projectSquadRepository) applyPreloads(db *gorm.DB) *gorm.DB {
	for _, relation := range repo.preloads {
		db = db.Preload(relation)
	}
	return db
}

func (repo *projectSquadRepository) InsertProjectSquad(data *models.ProjectSquad) (*models.ProjectSquad, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectSquadRepository) UpdateProjectSquad(id int64, updates map[string]interface{}) (*models.ProjectSquad, error) {
	if err := repo.db.Model(&models.ProjectSquad{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	var updated models.ProjectSquad
	if err := repo.db.First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (repo *projectSquadRepository) RemoveProjectSquad(data *models.ProjectSquad) (*models.ProjectSquad, error) {
	if err := repo.db.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectSquadRepository) RemoveProjectSquadByID(id int64) error {
	return repo.db.Delete(&models.ProjectSquad{}, id).Error
}

func (repo *projectSquadRepository) FindProjectSquad() ([]models.ProjectSquad, error) {
	var items []models.ProjectSquad
	db := repo.applyPreloads(repo.db.Model(&models.ProjectSquad{}))

	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *projectSquadRepository) FindProjectSquadByID(id int64) (*models.ProjectSquad, error) {
	var item models.ProjectSquad
	db := repo.applyPreloads(repo.db)

	if err := db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
