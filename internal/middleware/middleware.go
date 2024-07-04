package middleware

import (
	"net/http"
	"seen/internal/pkg"

	"github.com/gofiber/fiber/v2"
)

type (
	MiddlewareService interface {
		TokenAuthentication(*fiber.Ctx) error
	}
	middlewareService struct {
		jwtService pkg.JWTService
	}
)

func NewMIddlewareService(jwtService pkg.JWTService) MiddlewareService {
	return &middlewareService{
		jwtService: jwtService,
	}
}

// TokenAuthentication is a middleware for token validation.
func (c *middlewareService) TokenAuthentication(ctx *fiber.Ctx) error {
	// Extract the token from the "Authorization" header
	token := ctx.Get("Authorization")

	// Check if the token is empty or does not start with "Bearer "
	if token == "" || token[:7] != "Bearer " {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "خطا در اعتبار سنجی. کد خطا 55",
		})
	}

	// Extract the token string (remove "Bearer " prefix)
	token = token[7:]

	res, err := c.jwtService.VerifyToken(token)
	if err != nil {
		return ctx.Status(err.ErrorToHttpStatus()).JSON(err.ErrorToJsonMessage())
	}

	ctx.Locals("user_data", res.SID)

	// c.Locals("user_id", token)

	// Token is valid, continue with the next middleware
	return ctx.Next()
}
