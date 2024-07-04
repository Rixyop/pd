package services

import (
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
)

type (
	GarrisonService interface {
		GetGarrisonByGarrisonId(int32) (*models.Garrison, *types.Error)
	}
	garrisonService struct {
		repository repository.SeenRepository
	}
)

func NewGarrisonService(repository repository.SeenRepository) GarrisonService {
	return &garrisonService{
		repository: repository,
	}
}

func (c *garrisonService) GetGarrisonByGarrisonId(garrisonId int32) (*models.Garrison, *types.Error) {
	return c.repository.GetGarrisonByGarrisonId(garrisonId)
}
