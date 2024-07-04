package controllers

import (
	"seen/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	GarrisonController interface {
		GetGarrisonByGarrisonId(*fiber.Ctx) error
	}
	garrisonController struct {
		garrisonService services.GarrisonService
	}
)

func NewGarrisonController(garrisonService services.GarrisonService) GarrisonController {
	return &garrisonController{
		garrisonService: garrisonService,
	}
}

func (c *garrisonController) GetGarrisonByGarrisonId(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1012",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "آدرس جلسه نباید خالی باشد. کد خطا 1013",
			"success": false,
		})
	}

	res, rerr := c.garrisonService.GetGarrisonByGarrisonId(int32(garrisonId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}
