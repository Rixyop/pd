package services

import (
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
	"seen/internal/utils"
)

type (
	ClusterService interface {
		AddCluster(*models.ClusterDTO) (*models.Cluster, *types.Error)
		GetListOfClusterOfCompany(*models.ClsIds) ([]models.Cluster, *types.Error)
		DeleteClusterByGAndBAndCAndC(*models.ClsIds) *types.Error
	}
	clusterService struct {
		repository repository.SeenRepository
	}
)

func NewClusterService(repository repository.SeenRepository) ClusterService {
	return &clusterService{
		repository: repository,
	}
}

func (c *clusterService) AddCluster(data *models.ClusterDTO) (*models.Cluster, *types.Error) {

	clusterIdGen, rerr := utils.NextNumericString(40)
	if rerr != nil {
		return nil, types.NewBadRequestError("خطای داخلی رخ داده است. کد خطا 9")
	}
	data.Id = clusterIdGen
	_, err := c.repository.GetCompanyByGarrisonIdAndBattalionIdAndCompanyId(data.GarrisonId, data.BattalionId, data.CompanyId)
	if err != nil {
		return nil, err
	}
	return c.repository.AddCluster(data)
}

func (c *clusterService) GetListOfClusterOfCompany(data *models.ClsIds) ([]models.Cluster, *types.Error) {
	return c.repository.GetListOfClusterOfCompany(data)
}

func (c *clusterService) DeleteClusterByGAndBAndCAndC(data *models.ClsIds) *types.Error {
	return c.repository.DeleteClusterByGAndBAndCAndC(data)
}
