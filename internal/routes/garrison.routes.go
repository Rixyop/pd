package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func GarrisonGroup(app fiber.Router, garrionController controllers.GarrisonController, middleware middleware.MiddlewareService) {
	garrisonGroup := app.Group("/garrison")

	garrisonGroup.Use(middleware.TokenAuthentication)

	garrisonGroup.Get("/get/:garrison_id", garrionController.GetGarrisonByGarrisonId)
}
