package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func CourseCoachGroup(app fiber.Router, courseController controllers.CourseCoachController, middleware middleware.MiddlewareService) {
	courseCoachGroup := app.Group("/cc")

	courseCoachGroup.Use(middleware.TokenAuthentication)

	courseCoachGroup.Post("/add", courseController.AddCourseCoach)
	courseCoachGroup.Get("/get/list/:garrison_id", courseController.GetAllCourseCoachesByGarrisonId)
	courseCoachGroup.Delete("/del/:course_id", courseController.DeleteCourseCoachByCourseId)
}
