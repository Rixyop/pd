package controllers

import (
	"seen/internal/models"
	"seen/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	CourseController interface {
		AddCourse(*fiber.Ctx) error
		GetAllCoursesByGarrisonId(*fiber.Ctx) error
		DeleteCourseByGarrisonIdAndCourseId(*fiber.Ctx) error
	}
	courseController struct {
		courseService services.CourseService
	}
)

func NewCourseService(courseService services.CourseService) CourseController {
	return &courseController{
		courseService: courseService,
	}
}

func (c *courseController) AddCourse(ctx *fiber.Ctx) error {
	courseDTO := new(models.CourseDTO)
	if err := ctx.BodyParser(courseDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1030",
			"success": false,
		})
	}
	res, err := c.courseService.AddCourse(courseDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *courseController) GetAllCoursesByGarrisonId(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1031",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه جلسه نباید خالی باشد. کد خطا 1032",
			"success": false,
		})
	}

	res, rerr := c.courseService.GetAllCoursesByGarrisonId(int32(garrisonId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *courseController) DeleteCourseByGarrisonIdAndCourseId(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1031",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1032",
			"success": false,
		})
	}

	couseIdS := ctx.Params("course_id")
	if couseIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه درس نباید خالی باشد. کد خطا 1033",
			"success": false,
		})
	}

	rerr := c.courseService.DeleteCourseByGarrisonIdAndCourseId(int32(garrisonId), couseIdS)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "درس با موفقیت حذف شد",
		"success": true,
	})
}
