package services

import (
	"regexp"
	"seen/internal/models"
	"seen/internal/repository"
	"seen/internal/types"
	"seen/internal/utils"
)

type (
	UserService interface {
		GetUserInfoByUserId(string) (*models.User, *types.Error)
		GetGarrisonUsers(int32) ([]models.User, *types.Error)
		GetGarrisonUsersByRole(int32, string) ([]models.User, *types.Error)
		AddUser(*models.UserDTO) *types.Error
	}
	userService struct {
		repository repository.SeenRepository
	}
)

func NewUserService(repository repository.SeenRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (c *userService) GetUserInfoByUserId(userId string) (*models.User, *types.Error) {
	return c.repository.GetUserByUserId(userId)
}

func (c *userService) GetGarrisonUsers(garrisonId int32) ([]models.User, *types.Error) {
	return c.repository.GetGarrisonUsersByGarrisonId(garrisonId)
}

func (c *userService) GetGarrisonUsersByRole(garrisonId int32, role string) ([]models.User, *types.Error) {
	return c.repository.GetGarrisonUsersByGarrisonIdAndRole(garrisonId, role)
}

func (c *userService) AddUser(data *models.UserDTO) *types.Error {
	regex := regexp.MustCompile(`^[a-z]+(?:[._][a-z]+)*$`)
	if !regex.MatchString(data.Username) {
		return types.NewBadRequestError("نام کاربری نباید سمبل و عدد داشته باشد. کد خطا 4")
	}

	hashedPassword, err := utils.HashPassword([]byte(data.Password))
	if err != nil {
		return err
	}
	data.Password = hashedPassword
	userIdGen, rerr := utils.NextAlphanumericString(40)
	if rerr != nil {
		return types.NewBadRequestError("خطای داخلی رخ داده است. کد خطا 6")
	}
	data.UserId = userIdGen
	res, err := c.repository.GetUserByUsername(data.Username)
	if err != nil {
		if err.Code != 5 {
			return err
		}
	}
	if res != nil {
		return types.NewBadRequestError("نام کابری قبلا ثبت شده است. کد خطا 5")
	}

	err = c.repository.AddUser(data)
	if err != nil {
		return err
	}

	return c.repository.AddPassword(data.UserId, data.Password)
}
