package models

type PhotoSchema struct {
	// gorm.Model
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url"`
}
