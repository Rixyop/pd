package controllers

import (
	"seen/internal/models"
	"seen/internal/services"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type (
	UserController interface {
		GetSelfInfo(*fiber.Ctx) error
		GetUserInfoByUserId(*fiber.Ctx) error
		GetGarrisonUsers(*fiber.Ctx) error
	}
	userController struct {
		userService services.UserService
	}
)

func NewUserController(userService services.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) GetSelfInfo(ctx *fiber.Ctx) error {
	localData := ctx.Locals("user_data")
	if localData == nil {
		return ctx.JSON(map[string]interface{}{
			"message": "کاربر پیدا نشد. کد خطا 1001",
			"success": false,
		})
	}
	userData := localData.(models.Token)
	res, err := c.userService.GetUserInfoByUserId(userData.UserId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetUserInfoByUserId(ctx *fiber.Ctx) error {
	userId := ctx.Params("user_id")
	if userId == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": " شناسه کاربر نباید خالی باشد. کد خطا 1002",
			"success": false,
		})
	}
	res, err := c.userService.GetUserInfoByUserId(userId)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetGarrisonUsers(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1003",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "آدرس جلسه نباید خالی باشد. کد خطا 1004",
			"success": false,
		})
	}
	res, rerr := c.userService.GetGarrisonUsers(int32(garrisonId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) GetGarrisonUsersByRole(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1003",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "آدرس جلسه نباید خالی باشد. کد خطا 1004",
			"success": false,
		})
	}
	roleId := ctx.Params("role_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "نقش کاربر نباید خالی باشد. کد خطا 1005",
			"success": false,
		})
	}
	res, rerr := c.userService.GetGarrisonUsersByRole(int32(garrisonId), roleId)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *userController) AddUser(ctx *fiber.Ctx) error {
	userDTO := new(models.UserDTO)
	if err := ctx.BodyParser(userDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1006",
			"success": false,
		})
	}
	userDTO.CreatedAt = time.Now().Local()
	err := c.userService.AddUser(userDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "حساب کابری با موفقیت ساخته شد",
		"success": true,
	})
}
