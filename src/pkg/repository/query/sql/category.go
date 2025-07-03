
	package sql

	import (
		"jk-api/internal/database/config"
		"jk-api/internal/database/models"
		"jk-api/pkg/repository/adapter/sql"

		"gorm.io/gorm"
	)

	type categoryRepository struct {
		db       *gorm.DB
		preloads []string
	}

	func NewCategoryRepository() sql.CategoryRepository {
		return &categoryRepository{db: config.DB}
	}

	func (repo *categoryRepository) WithTx(tx *gorm.DB) sql.CategoryRepository {
		return &categoryRepository{
			db:       tx,
			preloads: repo.preloads,
		}
	}

	func (repo *categoryRepository) WithPreloads(preloads ...string) sql.CategoryRepository {
		return &categoryRepository{
			db:       repo.db,
			preloads: preloads,
		}
	}

	func (repo *categoryRepository) applyPreloads(db *gorm.DB) *gorm.DB {
		for _, relation := range repo.preloads {
			db = db.Preload(relation)
		}
		return db
	}

	func (repo *categoryRepository) InsertCategory(data *models.Category) (*models.Category, error) {
		if err := repo.db.Create(data).Error; err != nil {
			return nil, err
		}
		return data, nil
	}

	func (repo *categoryRepository) UpdateCategory(id int64, updates map[string]interface{}) (*models.Category, error) {
		if err := repo.db.Model(&models.Category{}).Where("id = ?", id).Updates(updates).Error; err != nil {
			return nil, err
		}

		var updated models.Category
		if err := repo.db.First(&updated, id).Error; err != nil {
			return nil, err
		}

		return &updated, nil
	}

	func (repo *categoryRepository) RemoveCategory(data *models.Category) (*models.Category, error) {
		if err := repo.db.Delete(data).Error; err != nil {
			return nil, err
		}
		return data, nil
	}

	func (repo *categoryRepository) RemoveCategoryByID(id int64) error {
		return repo.db.Delete(&models.Category{}, id).Error
	}

	func (repo *categoryRepository) FindCategory() ([]models.Category, error) {
		var items []models.Category
		db := repo.applyPreloads(repo.db.Model(&models.Category{}))

		if err := db.Find(&items).Error; err != nil {
			return nil, err
		}

		return items, nil
	}

	func (repo *categoryRepository) FindCategoryByID(id int64) (*models.Category, error) {
		var item models.Category
		db := repo.applyPreloads(repo.db)

		if err := db.First(&item, id).Error; err != nil {
			return nil, err
		}
		return &item, nil
	}
	