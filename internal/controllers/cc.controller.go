package controllers

import (
	"seen/internal/models"
	"seen/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	CourseCoachController interface {
		AddCourseCoach(*fiber.Ctx) error
		GetAllCourseCoachesByGarrisonId(*fiber.Ctx) error
		DeleteCourseCoachByCourseId(*fiber.Ctx) error
	}
	courseCoachController struct {
		courseCoachService services.CourseCoachService
	}
)

func NewCourseCoachController(courseCoachService services.CourseCoachService) CourseCoachController {
	return &courseCoachController{
		courseCoachService: courseCoachService,
	}
}

func (c *courseCoachController) AddCourseCoach(ctx *fiber.Ctx) error {
	courseCoachDTO := new(models.CourseCoachDTO)
	if err := ctx.BodyParser(courseCoachDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1034",
			"success": false,
		})
	}

	err := c.courseCoachService.AddCourseCoach(courseCoachDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	return ctx.JSON(map[string]interface{}{
		"message": "مربی/درس اضافه شد",
		"success": true,
	})
}

func (c *courseCoachController) GetAllCourseCoachesByGarrisonId(ctx *fiber.Ctx) error {
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
	res, rerr := c.courseCoachService.GetAllCourseCoachesByGarrisonId(int32(garrisonId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *courseCoachController) DeleteCourseCoachByCourseId(ctx *fiber.Ctx) error {
	courseIdS := ctx.Params("course_id")
	if courseIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه درس نباید خالی باشد. کد خطا 1031",
			"success": false,
		})
	}
	rerr := c.courseCoachService.DeleteCourseCoachByCourseId(courseIdS)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "درس/مربی با موفقیت حذف شد",
		"success": true,
	})
}
