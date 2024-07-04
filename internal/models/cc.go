package models

type CourseCoach struct {
	CourseId   string `json:"course_id"`
	CoachId    string `json:"coach_id"`
	GarrisonId int32  `json:"garrison_id"`
}
