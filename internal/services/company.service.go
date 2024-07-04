package services

import (
	"fmt"
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
	"seen/internal/utils"
)

type (
	CompanyService interface {
		AddCompany(*models.CompanyDTO) (*models.Company, *types.Error)
		GetListOfCompanyOfBattalion(int32, int32) ([]models.Company, *types.Error)
		DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(int32, int32, string) *types.Error
	}
	companyService struct {
		repository repository.SeenRepository
	}
)

func NewCompanyService(repository repository.SeenRepository) CompanyService {
	return &companyService{
		repository: repository,
	}
}

func (c *companyService) AddCompany(data *models.CompanyDTO) (*models.Company, *types.Error) {
	fmt.Println(data)

	companyIdGen, rerr := utils.NextNumericString(40)
	if rerr != nil {
		return nil, types.NewBadRequestError("خطای داخلی رخ داده است. کد خطا 8")
	}
	data.Id = companyIdGen
	_, err := c.repository.GetBattalionByGarrisonIdAndBattalionId(data.GarrisonId, data.BattalionId)
	if err != nil {
		return nil, err
	}
	return c.repository.AddCompany(data)
}

func (c *companyService) GetListOfCompanyOfBattalion(garrisonId int32, battalionId int32) ([]models.Company, *types.Error) {
	return c.repository.GetListOfCompanyOfBattalion(garrisonId, battalionId)
}

func (c *companyService) DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(garrisonId int32, battalionId int32, companyId string) *types.Error {
	return c.repository.DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(garrisonId, battalionId, companyId)
}
