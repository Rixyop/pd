package models

type GarrisonDTO struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	Creator   string
	CreatedAt string
}
