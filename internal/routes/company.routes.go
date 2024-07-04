package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func CompanyGroup(app fiber.Router, companyController controllers.CompanyController, middleware middleware.MiddlewareService) {
	companyGroup := app.Group("/company")

	companyGroup.Use(middleware.TokenAuthentication)

	companyGroup.Post("/add", companyController.AddCompany)
	companyGroup.Get("/get/list/:garrison_id/:battalion_id", companyController.GetListOfCompanyOfBattalion)
	companyGroup.Delete("/del/:garrison_id/:battalion_id/:company_id", companyController.DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId)
}
