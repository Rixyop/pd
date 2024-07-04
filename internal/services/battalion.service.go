package services

import (
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
)

type (
	BattalionService interface {
		AddBattalion(*models.BattalionDTO) (*models.Battalion, *types.Error)
		GetListOfBattalionOfGarrison(int32) ([]models.Battalion, *types.Error)
		DeleteBattalionByGarrisonIdAndBattalionId(int32, int32) *types.Error
	}
	battalionService struct {
		repository repository.SeenRepository
	}
)

func NewBattalionService(repository repository.SeenRepository) BattalionService {
	return &battalionService{
		repository: repository,
	}
}

func (c *battalionService) AddBattalion(data *models.BattalionDTO) (*models.Battalion, *types.Error) {
	// battalionIdGen, rerr := utils.NextNumericString(40)
	// if rerr != nil {
	// 	return nil, types.NewBadRequestError("خطای داخلی رخ داده است. کد خطا 7")
	// }
	// data.Id = battalionIdGen
	_, err := c.repository.GetGarrisonByGarrisonId(data.GarrisonId)
	if err != nil {
		return nil, err
	}
	return c.repository.AddBattalion(data)
}

func (c *battalionService) GetListOfBattalionOfGarrison(garrisonId int32) ([]models.Battalion, *types.Error) {
	return c.repository.GetListOfBattalionOfGarrison(garrisonId)
}

func (c *battalionService) DeleteBattalionByGarrisonIdAndBattalionId(garrisonId int32, battalionId int32) *types.Error {
	return c.repository.DeleteBattalionByGarrisonIdAndBattalionId(garrisonId, battalionId)
}
