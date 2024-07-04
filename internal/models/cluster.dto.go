package models

type ClusterDTO struct {
	Id          string
	GarrisonId  int32  `json:"garrison_id"`
	BattalionId int32  `json:"battalion_id"`
	CompanyId   string `json:"company_id"`
	Name        string `json:"name"`
}
