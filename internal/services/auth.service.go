package services

import (
	"fmt"
	"seen/internal/models"
	"seen/internal/pkg"
	"seen/internal/repository"
	"seen/internal/types"
	"seen/internal/utils"
)

type (
	AuthService interface {
		Login(*models.LoginDTO) (string, *types.Error)
	}
	authService struct {
		repository repository.SeenRepository
		jwtService pkg.JWTService
	}
)

func NewAuthService(repository repository.SeenRepository, jwtService pkg.JWTService) AuthService {
	return &authService{
		repository: repository,
		jwtService: jwtService,
	}
}

func (c *authService) Login(loginDTO *models.LoginDTO) (string, *types.Error) {
	userData, err := c.repository.GetUserByUsername(loginDTO.Username)
	if err != nil {
		if err.Code == 5 {
			return "", types.NewBadRequestError("رمز یا نام کاربری اشتباه است. کد خطا 3-2")

		} else {
			return "", err
		}
	}

	passwordData, err := c.repository.GetPasswordByUserId(userData.Id)
	if err != nil {
		return "", err
	}

	ok, err := utils.VerifyPassword([]byte(loginDTO.Password), []byte(passwordData))
	if err != nil {
		return "", err
	}
	if !ok {
		return "", types.NewBadRequestError("رمز یا نام کاربری اشتباه است. کد خطا 3")
	}
	var expireTime int32
	if loginDTO.RememberMe {
		expireTime = 360
	} else {
		expireTime = 60
	}
	token, err := c.jwtService.GenerateToken(models.Token{
		UserId:     userData.Id,
		Role:       userData.Role,
		GarrisonId: userData.GarrisonId,
	}, expireTime)
	if err != nil {
		return "", err
	}

	fmt.Println(token)
	return token, nil
}
