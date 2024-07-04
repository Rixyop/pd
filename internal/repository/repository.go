package repository

import (
	"database/sql"
	"seen/internal/models"
	"seen/internal/types"
)

type (
	SeenRepository interface {
		UserExistsByUserId(string) (bool, *types.Error)
		GetUserByUsername(string) (*models.User, *types.Error)
		GetUserByUserId(string) (*models.User, *types.Error)
		GetPasswordByUserId(string) (string, *types.Error)
		AddUser(*models.UserDTO) *types.Error
		AddPassword(string, string) *types.Error

		AddGarrison(*models.GarrisonDTO) (*models.Garrison, *types.Error)
		GetGarrisonUsersByGarrisonId(int32) ([]models.User, *types.Error)
		GetGarrisonUsersByGarrisonIdAndRole(int32, string) ([]models.User, *types.Error)
		GetGarrisonByGarrisonId(int32) (*models.Garrison, *types.Error)

		AddBattalion(*models.BattalionDTO) (*models.Battalion, *types.Error)
		GetListOfBattalionOfGarrison(int32) ([]models.Battalion, *types.Error)
		DeleteBattalionByGarrisonIdAndBattalionId(int32, int32) *types.Error
		GetBattalionByGarrisonIdAndBattalionId(int32, int32) (*models.Battalion, *types.Error)

		AddCompany(*models.CompanyDTO) (*models.Company, *types.Error)
		GetListOfCompanyOfBattalion(int32, int32) ([]models.Company, *types.Error)
		DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(int32, int32, string) *types.Error
		GetCompanyByGarrisonIdAndBattalionIdAndCompanyId(int32, int32, string) (*models.Company, *types.Error)

		AddCluster(*models.ClusterDTO) (*models.Cluster, *types.Error)
		GetListOfClusterOfCompany(*models.ClsIds) ([]models.Cluster, *types.Error)
		DeleteClusterByGAndBAndCAndC(*models.ClsIds) *types.Error

		AddCourse(*models.CourseDTO) (*models.Course, *types.Error)
		GetAllCoursesByGarrisonId(int32) ([]models.Course, *types.Error)
		DeleteCourseByGarrisonIdAndCourseId(int32, string) *types.Error
		CheckCourseByCourseId(string) (bool, *types.Error)

		DeleteCourseCoachByCourseId(string) *types.Error
		CheckCourseCoachByCourseIdAndCoachId(string, string) (bool, *types.Error)
		AddCourseCoach(string, string, int32) *types.Error
		GetAllCourseCoachesByGarrisonId(int32) ([]models.CourseCoach, *types.Error)
	}
	seenRepository struct {
		db *sql.DB
	}
)

func NewSeenRepository(db *sql.DB) SeenRepository {
	return &seenRepository{
		db: db,
	}
}
