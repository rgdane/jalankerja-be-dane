package container

import (
	"jk-api/api/controllers/v1/handlers"
)

type AppContainer struct {
	SquadHandler *handlers.SquadHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		SquadHandler: InitSquadContainer(),
	}
}
