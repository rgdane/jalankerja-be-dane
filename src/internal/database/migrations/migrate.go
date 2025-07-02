package migrations

import (
	"jk-api/internal/database/config"
	"jk-api/internal/database/models"
	"log"
)

func Migrate() {
	db := config.DB

	err := db.AutoMigrate(
		&models.User{},
		&models.Squad{},
		&models.Task{},
		&models.Project{},
		&models.ProjectSquad{},
	)

	if err != nil {
		log.Fatalf("❌ Migration failed: %v", err)
	}

	log.Println("✅ Migration complete")
}
