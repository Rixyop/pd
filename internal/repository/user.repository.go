package repository

import (
	"database/sql"
	"seen/internal/models"
	"seen/internal/types"
)

func (c *seenRepository) UserExistsByUserId(userId string) (bool, *types.Error) {
	var result int32
	err := c.db.QueryRow("SELECT 1 FROM users WHERE id = $1", userId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 100")
	}
	return true, nil
}

func (c *seenRepository) GetUserByUsername(username string) (*models.User, *types.Error) {
	query := `SELECT id, first_name, last_name, username, role, avatar, garrison_id, created_at FROM users WHERE username = $1`
	var user models.User
	err := c.db.QueryRow(query, username).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Role, &user.Avatar, &user.GarrisonId, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("کاربری پیدا نشد. کد خطا 101")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 102")
	}
	return &user, nil
}
func (c *seenRepository) GetUserByUserId(userId string) (*models.User, *types.Error) {
	query := `SELECT id, first_name, last_name, username, role, avatar, garrison_id, created_at FROM users WHERE id = $1`
	var user models.User
	err := c.db.QueryRow(query, userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Role, &user.Avatar, &user.GarrisonId, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("کاربری پیدا نشد. کد خطا 103")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 104")
	}
	return &user, nil
}

func (c *seenRepository) GetPasswordByUserId(userId string) (string, *types.Error) {
	query := `SELECT password FROM passwords WHERE user_id = $1`
	var password string
	err := c.db.QueryRow(query, userId).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", types.NewNotFoundError("کاربری پیدا نشد. کد خطا 105")
		}
		return "", types.NewInternalError("خطای داخلی رخ داده است. کد خطا 106")
	}
	return password, nil
}

func (c *seenRepository) AddUser(data *models.UserDTO) *types.Error {
	query := `INSERT INTO users (id, first_name, last_name, username, role, avatar, garrison_id, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`
	_, err := c.db.Exec(query, &data.UserId, &data.FirstName, &data.LastName, &data.Username, &data.Role, &data.Avatar, &data.GarrisonId, &data.CreatedAt)
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 115")
	}
	return nil
}

func (c *seenRepository) AddPassword(userId string, password string) *types.Error {
	query := `INSERT INTO passwords (user_id, password, last_update_at) VALUES ($1,$2,NOW())`
	_, err := c.db.Exec(query, &userId, &password)
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 116")
	}

	return nil
}
