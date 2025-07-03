package sql

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"jk-api/pkg/repository/adapter/sql"

	"gorm.io/gorm"
)

type projectRepository struct {
	db       *gorm.DB
	preloads []string
}

func NewProjectRepository() sql.ProjectRepository {
	return &projectRepository{db: config.DB}
}

func (repo *projectRepository) WithTx(tx *gorm.DB) sql.ProjectRepository {
	return &projectRepository{
		db:       tx,
		preloads: repo.preloads,
	}
}

func (repo *projectRepository) WithPreloads(preloads ...string) sql.ProjectRepository {
	return &projectRepository{
		db:       repo.db,
		preloads: preloads,
	}
}

func (repo *projectRepository) applyPreloads(db *gorm.DB) *gorm.DB {
	for _, relation := range repo.preloads {
		db = db.Preload(relation)
	}
	return db
}

func (repo *projectRepository) InsertProject(data *models.Project) (*models.Project, error) {
	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectRepository) UpdateProject(id int64, updates map[string]interface{}) (*models.Project, error) {
	if err := repo.db.Model(&models.Project{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}

	var updated models.Project
	if err := repo.db.First(&updated, id).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}

func (repo *projectRepository) RemoveProject(data *models.Project) (*models.Project, error) {
	if err := repo.db.Delete(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (repo *projectRepository) RemoveProjectByID(id int64) error {
	return repo.db.Delete(&models.Project{}, id).Error
}

func (repo *projectRepository) FindProject() ([]models.Project, error) {
	var items []models.Project
	db := repo.applyPreloads(repo.db.Model(&models.Project{}))

	if err := db.Order("created_at ASC").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (repo *projectRepository) FindProjectByID(id int64) (*models.Project, error) {
	var item models.Project
	db := repo.applyPreloads(repo.db)

	if err := db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
