package repository

import (
	"database/sql"
	"seen/internal/models"
	"seen/internal/types"
)

type (
	SeenRepository interface {
		UserExistsByUserId(string) (bool, *types.Error)
		GetUserByUsername(string) (*models.User, *types.Error)
		GetUserByUserId(string) (*models.User, *types.Error)
		GetPasswordByUserId(string) (string, *types.Error)
		GetGarrisonUsersByGarrisonId(int32) ([]models.User, *types.Error)
		GetGarrisonUsersByGarrisonIdAndRole(int32, string) ([]models.User, *types.Error)
		AddUser(*models.UserDTO) *types.Error
	}
	seenRepository struct {
		db *sql.DB
	}
)

func NewSeenRepository(db *sql.DB) SeenRepository {
	return &seenRepository{
		db: db,
	}
}
