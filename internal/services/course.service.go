package services

import (
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
	"seen/internal/utils"
)

type (
	CourseService interface {
		AddCourse(*models.CourseDTO) (*models.Course, *types.Error)
		GetAllCoursesByGarrisonId(int32) ([]models.Course, *types.Error)
		DeleteCourseByGarrisonIdAndCourseId(int32, string) *types.Error
	}
	courseService struct {
		repository repository.SeenRepository
	}
)

func NewCourseService(repository repository.SeenRepository) CourseService {
	return &courseService{
		repository: repository,
	}
}

func (c *courseService) AddCourse(data *models.CourseDTO) (*models.Course, *types.Error) {
	courseIdGen, rerr := utils.NextNumericString(40)
	if rerr != nil {
		return nil, types.NewBadRequestError("خطای داخلی رخ داده است. کد خطا 10")
	}
	data.Id = courseIdGen
	_, err := c.repository.GetGarrisonByGarrisonId(data.GarrisonId)
	if err != nil {
		return nil, err
	}
	return c.repository.AddCourse(data)
}

func (c *courseService) GetAllCoursesByGarrisonId(garrisonId int32) ([]models.Course, *types.Error) {
	return c.repository.GetAllCoursesByGarrisonId(garrisonId)
}

func (c *courseService) DeleteCourseByGarrisonIdAndCourseId(garrisonId int32, courseId string) *types.Error {
	err := c.repository.DeleteCourseCoachByCourseId(courseId)
	if err != nil {
		return err
	}
	return c.repository.DeleteCourseByGarrisonIdAndCourseId(garrisonId, courseId)
}
