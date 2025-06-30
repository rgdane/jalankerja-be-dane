package container

import (
	"jk-api/api/controllers/v1"
	"jk-api/api/controllers/v1/handlers"
	"jk-api/pkg/repository/query/sql"
	"jk-api/pkg/services/v1"
)

type ServiceContainer struct {
	SquadService services.SquadService
}

func NewServiceContainer() *ServiceContainer {
	return &ServiceContainer{
		SquadService: services.NewSquadService(sql.NewSquadRepository()),
	}
}

func (c *ServiceContainer) RegisterControllers() {
	controllers.SquadService = c.SquadService
	handlers.SquadService = c.SquadService
}
