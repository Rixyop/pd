package models

import "time"

type Battalion struct {
	Id         int32     `json:"id"`
	GarrisonId int32     `json:"garrison_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
}
