package models

type Token struct {
	UserId     string `json:"user_id"`
	Role       string `json:"role"`
	GarrisonId *int32 `json:"garrison_id"`
	// AccessToken string    `json:"access_token"`
	// ExpireAt    time.Time `json:"expire_at"`
	// CreatedAt time.Time `json:"created_at"`
}
