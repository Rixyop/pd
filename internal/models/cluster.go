package models

import "time"

type Cluster struct {
	Id          string    `json:"id"`
	GarrisonId  int32     `json:"garrison_id"`
	BattalionId int32     `json:"battalion_id"`
	CompanyId   string    `json:"company_id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
}
