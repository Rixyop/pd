package controllers

import (
	"seen/internal/models"
	"seen/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	CompanyController interface {
		AddCompany(*fiber.Ctx) error
		GetListOfCompanyOfBattalion(*fiber.Ctx) error
		DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(*fiber.Ctx) error
	}
	companyController struct {
		companyService services.CompanyService
	}
)

func NewCompanyController(companyService services.CompanyService) CompanyController {
	return &companyController{
		companyService: companyService,
	}
}

func (c *companyController) AddCompany(ctx *fiber.Ctx) error {
	companyDTO := new(models.CompanyDTO)
	if err := ctx.BodyParser(companyDTO); err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "خطا در درخواست. کد خطا 1018",
			"success": false,
		})
	}
	res, err := c.companyService.AddCompany(companyDTO)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *companyController) GetListOfCompanyOfBattalion(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1019",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "آدرس جلسه نباید خالی باشد. کد خطا 1020",
			"success": false,
		})
	}
	battalionIdS := ctx.Params("battalion_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گردان نباید خالی باشد. کد خطا 1021",
			"success": false,
		})
	}
	battalionId, err := strconv.ParseInt(battalionIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گردان نباید خالی باشد. کد خطا 1022",
			"success": false,
		})
	}

	res, rerr := c.companyService.GetListOfCompanyOfBattalion(int32(garrisonId), int32(battalionId))
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"data":    res,
		"success": true,
	})
}

func (c *companyController) DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(ctx *fiber.Ctx) error {
	garrisonIdS := ctx.Params("garrison_id")
	if garrisonIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1023",
			"success": false,
		})
	}
	garrisonId, err := strconv.ParseInt(garrisonIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه پادگان نباید خالی باشد. کد خطا 1024",
			"success": false,
		})
	}

	battalionIdS := ctx.Params("battalion_id")
	if battalionIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گردان نباید خالی باشد. کد خطا 1025",
			"success": false,
		})
	}
	battalionId, err := strconv.ParseInt(battalionIdS, 10, 32)
	if err != nil {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گردان نباید خالی باشد. کد خطا 1026",
			"success": false,
		})
	}

	companyIdS := ctx.Params("company_id")
	if companyIdS == "" {
		return ctx.Status(400).JSON(map[string]interface{}{
			"message": "شناسه گروهان نباید خالی باشد. کد خطا 1027",
			"success": false,
		})
	}

	rerr := c.companyService.DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(int32(garrisonId), int32(battalionId), companyIdS)
	if rerr != nil {
		return ctx.Status(rerr.ErrorToHttpStatus()).JSON(rerr.ErrorToJsonMessage())
	}
	return ctx.JSON(map[string]interface{}{
		"message": "گروهان با موفقیت حذف شد",
		"success": true,
	})
}
