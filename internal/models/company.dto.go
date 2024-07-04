package models

type CompanyDTO struct {
	Id          string
	GarrisonId  int32  `json:"garrison_id"`
	BattalionId int32  `json:"battalion_id"`
	Name        string `json:"name"`
}
