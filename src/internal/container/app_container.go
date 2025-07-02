package container

import (
	"jk-api/api/controllers/v1/handlers"
)

type AppContainer struct {
	ProjectHandler *handlers.ProjectHandler
	SquadHandler   *handlers.SquadHandler
	UserHandler    *handlers.UserHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		ProjectHandler: InitProjectContainer(),
		SquadHandler:   InitSquadContainer(),
		UserHandler:    InitUserContainer(),
	}
}
