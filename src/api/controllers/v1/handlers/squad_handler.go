package handlers

import (
	"jk-api/api/controllers/v1/dto"
	"jk-api/api/controllers/v1/mapper"
	"jk-api/internal/database/models"
	"jk-api/pkg/services/v1"
)

type SquadHandler struct {
	Service             services.SquadService
	ProjectSquadService services.ProjectSquadService
}

func NewSquadHandler(
	service services.SquadService,
	projectSquadService services.ProjectSquadService,
) *SquadHandler {
	return &SquadHandler{
		Service:             service,
		ProjectSquadService: projectSquadService,
	}
}

func (h *SquadHandler) CreateSquadHandler(input *dto.CreateSquadDto) (*dto.SquadResponseDto, error) {
	db := h.Service.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil || db.Error != nil {
			db.Rollback()
		}
	}()

	squadService := h.Service.WithTx(db)
	projectSquadService := h.ProjectSquadService.WithTx(db)

	createSquad, err := mapper.CreateSquadDtoToModel(input)
	if err != nil {
		return nil, err
	}

	createdData, err := squadService.CreateSquad(createSquad)
	if err != nil {
		return nil, err
	}

	if input.ProjectID != nil {
		projectSquad := &models.ProjectSquad{
			ProjectID: *input.ProjectID,
			SquadID:   createdData.ID,
		}
		if _, err := projectSquadService.CreateProjectSquad(projectSquad); err != nil {
			return nil, err
		}
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}

	return mapper.SquadModelToResponseDto(createdData)
}

func (h *SquadHandler) UpdateSquadHandler(id int64, input *dto.UpdateSquadDto) (*models.Squad, error) {
	db := h.Service.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			db.Rollback()
		}
	}()

	squadService := h.Service.WithTx(db)

	payload, err := mapper.UpdateSquadDtoToModel(input)
	if err != nil {
		return nil, err
	}

	updatedData, err := squadService.UpdateSquad(id, payload)
	if err != nil {
		return nil, err
	}

	if err := db.Commit().Error; err != nil {
		return nil, err
	}

	return updatedData, nil
}

func (h *SquadHandler) DeleteSquadHandler(id int64) error {
	return h.Service.DeleteSquad(id)
}

func (h *SquadHandler) GetSquadByIDHandler(id int64) (*models.Squad, error) {
	return h.Service.GetSquadByID(id)
}

func (h *SquadHandler) GetAllSquadsHandler() ([]models.Squad, error) {
	return h.Service.GetAllSquads()
}
