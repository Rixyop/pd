package models

type CourseDTO struct {
	Id         string
	GarrisonId int32  `json:"garrison_id"`
	Name       string `json:"name"`
	Count      int32  `json:"count"`
	Priority   int32  `json:"priority"`
	ClassTime  int32  `json:"class_time"`
}
