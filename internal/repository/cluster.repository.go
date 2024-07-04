package repository

import (
	"database/sql"
	"fmt"
	"seen/internal/models"
	"seen/internal/types"
	"time"
)

func (c *seenRepository) AddCluster(data *models.ClusterDTO) (*models.Cluster, *types.Error) {
	query := `INSERT INTO clusters (id, garrison_id, battalion_id, company_id, name, created_at) VALUES ($1,$2,$3,$4,$5,NOW())`
	_, err := c.db.Exec(query, &data.Id, &data.GarrisonId, &data.BattalionId, &data.CompanyId, &data.Name)
	if err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 136")
	}
	return &models.Cluster{
		Id:          data.Id,
		GarrisonId:  data.GarrisonId,
		BattalionId: data.BattalionId,
		CompanyId:   data.CompanyId,
		Name:        data.Name,
		CreatedAt:   time.Now(),
	}, nil
}

func (c *seenRepository) GetListOfClusterOfCompany(data *models.ClsIds) ([]models.Cluster, *types.Error) {
	query := `SELECT id, garrison_id, battalion_id, company_id, name, created_at FROM clusters WHERE garrison_id = $1 AND battalion_id = $2 AND company_id = $3`
	rows, err := c.db.Query(query, data.GarrisonId, data.BattalionId, data.CompanyId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("دسته ای پیدا نشد. کد خطا 137")
		}
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 138")
	}

	defer rows.Close()
	clusters := make([]models.Cluster, 0)
	for rows.Next() {
		var cluster models.Cluster
		if err := rows.Scan(
			&cluster.Id,
			&cluster.GarrisonId,
			&cluster.BattalionId,
			&cluster.CompanyId,
			&cluster.Name,
			&cluster.CreatedAt); err != nil {
			fmt.Println(err)
			return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 139")
		}
		clusters = append(clusters, cluster)
	}
	if err := rows.Err(); err != nil {
		return nil, types.NewInternalError("خطای داخلی رخ داده است. کد خطا 140")
	}
	return clusters, nil
}

func (c *seenRepository) DeleteClusterByGAndBAndCAndC(data *models.ClsIds) *types.Error {
	tx, err := c.db.Begin()
	if err != nil {
		return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 141")
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
		`DELETE FROM clusters WHERE id = $1 AND company_id = $2 AND battalion_id = $3 AND garrison_id = $4`,
	}

	for _, query := range queries {
		if _, err := tx.Exec(query, data.ClusterId, data.CompanyId, data.BattalionId, data.GarrisonId); err != nil {
			fmt.Println(err)
			return types.NewInternalError("خطای داخلی رخ داده است. کد خطا 142")
		}
	}

	return nil
}
