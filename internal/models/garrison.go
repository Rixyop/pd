package models

import "time"

type Garrison struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Location  string    `json:"location"`
	Creator   string    `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
}
