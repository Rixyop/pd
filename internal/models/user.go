package models

import "time"

type User struct {
	Id         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	Avatar     string    `json:"avatar"`
	GarrisonId *int32    `json:"garrison_id"`
	CreatedAt  time.Time `json:"created_at"`
}
