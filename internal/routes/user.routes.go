package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserGroup(app fiber.Router, userController controllers.UserController, middleware middleware.MiddlewareService) {
	userGroup := app.Group("/usr")

	userGroup.Use(middleware.TokenAuthentication)

	userGroup.Get("/get/self", userController.GetSelfInfo)
	userGroup.Get("/get/:user_id", userController.GetUserInfoByUserId)
	userGroup.Get("/get/garrison/:garrison_id", userController.GetGarrisonUsers)
	userGroup.Get("/get/garrison/:garrison_id/:role_id", userController.GetGarrisonUsersByRole)
	userGroup.Post("/add", userController.AddUser)
}
