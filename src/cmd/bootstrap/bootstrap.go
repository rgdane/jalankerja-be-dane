package bootstrap

import (
	"jk-api/internal/container"
	"jk-api/internal/database/config"
	"jk-api/internal/database/migrations"
	"log"
)

func InitApp() *container.AppContainer {
	if !config.PostgresInit() {
		log.Fatal("❌ Failed to connect to Postgres SQL")
	}

	services := container.NewAppContainer()
	migrations.Migrate()

	return services
}
