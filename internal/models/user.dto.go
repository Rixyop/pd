package models

import "time"

type LoginDTO struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type UserDTO struct {
	UserId     string
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Username   string `json:"username"`
	Role       string `json:"role"`
	Avatar     string `json:"avatar"`
	GarrisonId *int32 `json:"garrison_id"`
	Password   string `json:"password"`
	CreatedAt  time.Time
}
