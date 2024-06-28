package routes

import (
	"seen/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthGroup(app fiber.Router, authController controllers.AuthController) {
	// Create a new group for room-related routes.
	authGroup := app.Group("/auth")
	// roomGroup.Use(middleware.TokenAuthentication)
	authGroup.Post("/login", authController.Login)
}
