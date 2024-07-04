package models

type Course struct {
	Id         string `json:"id"`
	GarrisonId int32  `json:"garrison_id"`
	Name       string `json:"name"`
	Count      int32  `json:"count"`
	Priority   int32  `json:"priority"`
	ClassTime  int32  `json:"class_time"`
}
