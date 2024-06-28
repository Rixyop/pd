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
	query := `SELECT id, first_name, last_name, username, role, garrison_id, created_at FROM users WHERE username = $1`
	var user models.User
	err := c.db.QueryRow(query, username).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Role, &user.GarrisonId, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("کاربری پیدا نشد. کد خطا 101")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 102")
	}
	return &user, nil
}
func (c *seenRepository) GetUserByUserId(userId string) (*models.User, *types.Error) {
	query := `SELECT id, first_name, last_name, username, role, garrison_id, created_at FROM users WHERE id = $1`
	var user models.User
	err := c.db.QueryRow(query, userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Role, &user.GarrisonId, &user.CreatedAt)
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

func (c *seenRepository) GetGarrisonUsersByGarrisonId(garrisonId int32) ([]models.User, *types.Error) {
	query := `SELECT id, first_name, last_name, username, role, garrison_id, created_at FROM users WHERE garrison_id = $1`

	rows, err := c.db.Query(query, garrisonId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("کاربری پیدا نشد. کد خطا 107")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 108")
	}

	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Role,
			&user.GarrisonId,
			&user.CreatedAt); err != nil {
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 109")
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 110")
	}
	return users, nil
}

func (c *seenRepository) GetGarrisonUsersByGarrisonIdAndRole(garrisonId int32, role string) ([]models.User, *types.Error) {
	query := `SELECT id, first_name, last_name, username, role, garrison_id, created_at FROM users WHERE garrison_id = $1 AND role = $2`

	rows, err := c.db.Query(query, garrisonId, role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("کاربری پیدا نشد. کد خطا 111")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 112")
	}

	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Username,
			&user.Role,
			&user.GarrisonId,
			&user.CreatedAt); err != nil {
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 113")
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 114")
	}
	return users, nil
}

func (c *seenRepository) AddUser(data *models.UserDTO) *types.Error {
	query := `INSERT INTO users (first_name, last_name, username, role, avatar, garrison_id, created_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	_, err := c.db.Exec(query, &data.FirstName, &data.LastName, &data.Username, &data.Role, &data.Avatar, &data.GarrisonId, &data.CreatedAt)
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 115")
	}
	return nil
}
