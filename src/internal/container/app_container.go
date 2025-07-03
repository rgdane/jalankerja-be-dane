package container

import (
	"jk-api/api/controllers/v1/handlers"
)

type AppContainer struct {
	CategoryHandler *handlers.CategoryHandler
	ProjectHandler *handlers.ProjectHandler
	SquadHandler *handlers.SquadHandler
	UserHandler *handlers.UserHandler
}

func NewAppContainer() *AppContainer {
	return &AppContainer{
		CategoryHandler: InitCategoryContainer(),
		ProjectHandler: InitProjectContainer(),
		SquadHandler: InitSquadContainer(),
		UserHandler: InitUserContainer(),
	}
}
