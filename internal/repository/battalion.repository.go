package repository

import (
	"database/sql"
	"fmt"
	"seen/internal/models"
	"seen/internal/types"
	"time"
)

func (c *seenRepository) AddBattalion(data *models.BattalionDTO) (*models.Battalion, *types.Error) {
	query := `INSERT INTO battalions (garrison_id, name, created_at) VALUES ($1,$2,NOW()) RETURNING id`
	var battalionId int32
	err := c.db.QueryRow(query, &data.GarrisonId, &data.Name).Scan(&battalionId)
	if err != nil {
		fmt.Println(err)
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 117")
	}
	return &models.Battalion{
		Id:         battalionId,
		GarrisonId: data.GarrisonId,
		Name:       data.Name,
		CreatedAt:  time.Now(),
	}, nil
}

func (c *seenRepository) GetBattalionByGarrisonIdAndBattalionId(garrisonId int32, battalionId int32) (*models.Battalion, *types.Error) {
	query := `SELECT id, garrison_id, name, created_at FROM battalions WHERE garrison_id = $1 AND id = $2`

	var battalion models.Battalion

	err := c.db.QueryRow(query, garrisonId, battalionId).Scan(&battalion.Id, &battalion.GarrisonId, &battalion.Name, &battalion.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("گردانی پیدا نشد. کد خطا 134")
		}
		fmt.Println(err)
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 135")
	}

	return &battalion, nil
}

func (c *seenRepository) GetListOfBattalionOfGarrison(garrisonId int32) ([]models.Battalion, *types.Error) {
	query := `SELECT id, garrison_id, name, created_at FROM battalions WHERE garrison_id = $1`
	rows, err := c.db.Query(query, garrisonId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("گردانی پیدا نشد. کد خطا 119")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 120")
	}

	defer rows.Close()
	battalions := make([]models.Battalion, 0)
	for rows.Next() {
		var battalion models.Battalion
		if err := rows.Scan(
			&battalion.Id,
			&battalion.GarrisonId,
			&battalion.Name,
			&battalion.CreatedAt); err != nil {
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 121")
		}
		battalions = append(battalions, battalion)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 122")
	}
	return battalions, nil
}

func (c *seenRepository) DeleteBattalionByGarrisonIdAndBattalionId(garrisonId int32, battalionId int32) *types.Error {
	tx, err := c.db.Begin()
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 125")
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r) // re-throw panic after Rollback
		} else if err != nil {
			tx.Rollback() // err is non-nil; don't change it
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()

	queries := []string{
		`DELETE FROM clusters WHERE battalion_id = $1 AND garrison_id = $2`,
		`DELETE FROM companies WHERE battalion_id = $1 AND garrison_id = $2`,
		`DELETE FROM battalions WHERE id = $1 AND garrison_id = $2 AND garrison_id = $2`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query, battalionId, garrisonId); err != nil {
			fmt.Println(err)
			return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 126")
		}
	}

	return nil
}
