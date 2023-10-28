package app

type PhotoCreateInput struct {
	Title     string `json:"title" validate:"omitempty,max=20"`
	Caption   string `json:"caption" validate:"omitempty,max=60"`
	PhotoURL  string `json:"photo_url" validate:"required,url"`
	UserEmail string `json:"user_email" validate:"required,email"`
}

type PhotoUpdateInput struct {
	Title    string `json:"title" validate:"omitempty,max=20"`
	Caption  string `json:"caption" validate:"omitempty,max=60"`
	PhotoURL string `json:"photo_url" validate:"omitempty,url"`
}
