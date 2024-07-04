package controllers

import (
	"seen/internal/models"
	"seen/internal/services"

	"github.com/gofiber/fiber/v2"
)

type (
	ClusterController interface {
		AddCluster(*fiber.Ctx) error
		GetListOfClusterOfCompany(*fiber.Ctx) error
		DeleteClusterByGAndBAndCAndC(*fiber.Ctx) error
	}
	clusterController struct {
		clusterService services.ClusterService
	}
)

func NewClusterController(clusterService services.ClusterService) ClusterController {
	return &clusterController{
		clusterService: clusterService,
	}
}

func (c *clusterController) AddCluster(ctx *fiber.Ctx) error {
	clusterDTO := new(models.ClusterDTO)
	if err := ctx.BodyParser(clusterDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1028",
			"success": false,
		})
	}
	res, err := c.clusterService.AddCluster(clusterDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *clusterController) GetListOfClusterOfCompany(ctx *fiber.Ctx) error {
	clsDTO := new(models.ClsIds)
	if err := ctx.BodyParser(clsDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1028",
			"success": false,
		})
	}

	res, rerr := c.clusterService.GetListOfClusterOfCompany(clsDTO)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *clusterController) DeleteClusterByGAndBAndCAndC(ctx *fiber.Ctx) error {
	clsDTO := new(models.ClsIds)
	if err := ctx.BodyParser(clsDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1029",
			"success": false,
		})
	}

	rerr := c.clusterService.DeleteClusterByGAndBAndCAndC(clsDTO)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "دسته با موفقیت حذف شد",
		"success": true,
	})
}
