package routes

import (
	"seen/internal/controllers"
	"seen/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func ClusterGroup(app fiber.Router, clusterController controllers.ClusterController, middleware middleware.MiddlewareService) {
	clusterGroup := app.Group("/cluster")

	clusterGroup.Use(middleware.TokenAuthentication)

	clusterGroup.Post("/add", clusterController.AddCluster)
	clusterGroup.Post("/get/list", clusterController.GetListOfClusterOfCompany)
	clusterGroup.Post("/del", clusterController.DeleteClusterByGAndBAndCAndC)
}
