package controllers

import (
	"seen/internal/models"
	"seen/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	AuthController interface {
		Login(*fiber.Ctx) error
	}
	authController struct {
		authService services.AuthService
	}
)

func NewAuthController(authService services.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Login(ctx *fiber.Ctx) error {
	loginDTO := new(models.LoginDTO)
	if err := ctx.BodyParser(loginDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1000",
			"success": false,
		})
	}
	res, err := c.authService.Login(loginDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}
