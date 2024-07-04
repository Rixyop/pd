package repository

import (
	"database/sql"
	"fmt"
	"seen/internal/models"
	"seen/internal/types"
	"time"
)

func (c *seenRepository) AddCompany(data *models.CompanyDTO) (*models.Company, *types.Error) {
	query := `INSERT INTO companies (id, garrison_id, battalion_id, name, created_at) VALUES ($1,$2,$3,$4,NOW())`
	_, err := c.db.Exec(query, &data.Id, &data.GarrisonId, &data.BattalionId, &data.Name)
	if err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 127")
	}
	return &models.Company{
		Id:          data.Id,
		GarrisonId:  data.GarrisonId,
		BattalionId: data.BattalionId,
		Name:        data.Name,
		CreatedAt:   time.Now(),
	}, nil
}

func (c *seenRepository) GetListOfCompanyOfBattalion(garrisonId int32, battalionId int32) ([]models.Company, *types.Error) {
	query := `SELECT id, garrison_id, battalion_id, name, created_at FROM companies WHERE garrison_id = $1 AND battalion_id = $2`
	rows, err := c.db.Query(query, garrisonId, battalionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("گروهانی پیدا نشد. کد خطا 128")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 129")
	}

	defer rows.Close()
	companies := make([]models.Company, 0)
	for rows.Next() {
		var company models.Company
		if err := rows.Scan(
			&company.Id,
			&company.GarrisonId,
			&company.BattalionId,
			&company.Name,
			&company.CreatedAt); err != nil {
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 130")
		}
		companies = append(companies, company)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 131")
	}
	return companies, nil
}

func (c *seenRepository) DeleteCompanyByGarrisonIdAndBattalionIdAndCompanyId(garrisonId int32, battalionId int32, companyId string) *types.Error {
	tx, err := c.db.Begin()
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 132")
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
		`DELETE FROM clusters WHERE company_id = $1 AND battalion_id = $2 AND garrison_id = $3`,
		`DELETE FROM companies WHERE id = $1 AND battalion_id = $2 AND garrison_id = $3`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query, companyId, battalionId, garrisonId); err != nil {
			fmt.Println(err)
			return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 133")
		}
	}

	return nil
}

func (c *seenRepository) GetCompanyByGarrisonIdAndBattalionIdAndCompanyId(garrisonId int32, battalionId int32, companyId string) (*models.Company, *types.Error) {
	query := `SELECT id, garrison_id, battalion_id, name, created_at FROM companies WHERE garrison_id = $1 AND battalion_id = $2 AND id = $3`

	var company models.Company

	err := c.db.QueryRow(query, garrisonId, battalionId, companyId).Scan(&company.Id, &company.GarrisonId, &company.BattalionId, &company.Name, &company.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("گردانی پیدا نشد. کد خطا 143")
		}
		fmt.Println(err)
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 144")
	}

	return &company, nil
}
