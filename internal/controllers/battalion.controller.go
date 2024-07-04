package controllers

import (
	"seen/internal/models"
	"seen/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	BattalionController interface {
		AddBattalion(*fiber.Ctx) error
		GetListOfBattalionOfGarrison(*fiber.Ctx) error
		DeleteBattalionByGarrisonIdAndBattalionId(*fiber.Ctx) error
	}
	battalionController struct {
		battalionService services.BattalionService
	}
)

func NewBattalionController(battalionService services.BattalionService) BattalionController {
	return &battalionController{
		battalionService: battalionService,
	}
}

func (c *battalionController) AddBattalion(ctx *fiber.Ctx) error {
	battalionDTO := new(models.BattalionDTO)
	if err := ctx.BodyParser(battalionDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1007",
			"success": false,
		})
	}
	res, err := c.battalionService.AddBattalion(battalionDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *battalionController) GetListOfBattalionOfGarrison(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1008",
			"success": false,
		})
	}

	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "آدرس جلسه نباید خالی باشد. کد خطا 1009",
			"success": false,
		})
	}
	res, rerr := c.battalionService.GetListOfBattalionOfGarrison(int32(garrisonId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *battalionController) DeleteBattalionByGarrisonIdAndBattalionId(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1014",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه جلسه نباید خالی باشد. کد خطا 1015",
			"success": false,
		})
	}

	battalionIdS := ctx.Params("battalion_id")
	if battalionIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گردان نباید خالی باشد. کد خطا 1016",
			"success": false,
		})
	}
	battalionId, err := strconv.ParseInt(battalionIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گردان نباید خالی باشد. کد خطا 1017",
			"success": false,
		})
	}
	rerr := c.battalionService.DeleteBattalionByGarrisonIdAndBattalionId(int32(garrisonId), int32(battalionId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "گردان با موفقیت حذف شد",
		"success": true,
	})
}
