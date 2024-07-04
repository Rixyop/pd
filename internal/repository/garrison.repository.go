package repository

import (
	"database/sql"
	"seen/internal/models"
	"seen/internal/types"
	"time"
)

func (c *seenRepository) AddGarrison(data *models.GarrisonDTO) (*models.Garrison, *types.Error) {
	query := `INSERT INTO battalions (name, location, creator, created_at) VALUES ($1,$2,$3,NOW()) RETURNING "id"`
	var garrisonId int32
	err := c.db.QueryRow(query, &data.Name, &data.Location, &data.Creator).Scan(&garrisonId)
	if err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 118")
	}
	return &models.Garrison{
		Id:        garrisonId,
		Name:      data.Name,
		Location:  data.Location,
		Creator:   data.Creator,
		CreatedAt: time.Now(),
	}, nil
}

func (c *seenRepository) GetGarrisonUsersByGarrisonId(garrisonId int32) ([]models.User, *types.Error) {
	query := `SELECT id, first_name, last_name, username, role, avatar, garrison_id, created_at FROM users WHERE garrison_id = $1`

	rows, err := c.db.Query(query, garrisonId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("پادکانی پیدا نشد. کد خطا 107")
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
			&user.Avatar,
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
	query := `SELECT id, first_name, last_name, username, role, avatar, garrison_id, created_at FROM users WHERE garrison_id = $1 AND role = $2`

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
			&user.Avatar,
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

func (c *seenRepository) GetGarrisonByGarrisonId(garrisonId int32) (*models.Garrison, *types.Error) {
	query := `SELECT id, name, location, creator, created_at FROM garrisons WHERE id = $1`
	var garrison models.Garrison
	err := c.db.QueryRow(query, garrisonId).Scan(&garrison.Id, &garrison.Name, &garrison.Location, &garrison.Creator, &garrison.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("پادگانی پیدا نشد. کد خطا 123")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 124")
	}
	return &garrison, nil
}
