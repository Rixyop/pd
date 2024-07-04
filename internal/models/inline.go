package models

type IntIds struct {
	Id1 int32
	Id2 int32
	Id3 int32
	Id4 int32
}

type ClsIds struct {
	GarrisonId  int32  `json:"garrison_id"`
	BattalionId int32  `json:"battalion_id"`
	CompanyId   string `json:"company_id"`
	ClusterId   string `json:"cluster_id"`
}
