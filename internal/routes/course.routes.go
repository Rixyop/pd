package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func CourseGroup(app fiber.Router, courseController controllers.CourseController, middleware middleware.MiddlewareService) {
	courseGroup := app.Group("/course")

	courseGroup.Use(middleware.TokenAuthentication)

	courseGroup.Post("/add", courseController.AddCourse)
	courseGroup.Get("/get/list/:garrison_id", courseController.GetAllCoursesByGarrisonId)
	courseGroup.Delete("/del/:garrison_id/:course_id", courseController.DeleteCourseByGarrisonIdAndCourseId)
}
