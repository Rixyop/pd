package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func BattalionGroup(app fiber.Router, battalionController controllers.BattalionController, middleware middleware.MiddlewareService) {
	// Create a new group for room-related routes.
	battalionGroup := app.Group("/battalion")

	battalionGroup.Use(middleware.TokenAuthentication)
	// roomGroup.Use(middleware.TokenAuthentication)
	battalionGroup.Post("/add", battalionController.AddBattalion)
	battalionGroup.Get("/get/list/:garrison_id", battalionController.GetListOfBattalionOfGarrison)
	battalionGroup.Delete("/del/:garrison_id/:battalion_id", battalionController.DeleteBattalionByGarrisonIdAndBattalionId)
}
