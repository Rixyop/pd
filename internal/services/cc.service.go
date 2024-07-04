package services

import (
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
)

type (
	CourseCoachService interface {
		AddCourseCoach(*models.CourseCoachDTO) *types.Error
		GetAllCourseCoachesByGarrisonId(int32) ([]models.CourseCoach, *types.Error)
		DeleteCourseCoachByCourseId(string) *types.Error
	}
	courseCoachService struct {
		repository repository.SeenRepository
	}
)

func NewCourseCoachService(repository repository.SeenRepository) CourseCoachService {
	return &courseCoachService{
		repository: repository,
	}
}

func (c *courseCoachService) AddCourseCoach(data *models.CourseCoachDTO) *types.Error {
	_, err := c.repository.GetGarrisonByGarrisonId(data.GarrisonId)
	if err != nil {
		return err
	}

	_, err = c.repository.GetUserByUserId(data.CoachId)
	if err != nil {
		return err
	}

	_, err = c.repository.CheckCourseByCourseId(data.CoachId)
	if err != nil {
		return err
	}
	return c.repository.AddCourseCoach(data.CourseId, data.CoachId, data.GarrisonId)
}

func (c *courseCoachService) GetAllCourseCoachesByGarrisonId(garrisonId int32) ([]models.CourseCoach, *types.Error) {
	return c.repository.GetAllCourseCoachesByGarrisonId(garrisonId)
}

func (c *courseCoachService) DeleteCourseCoachByCourseId(courseId string) *types.Error {
	return c.repository.DeleteCourseCoachByCourseId(courseId)
}
