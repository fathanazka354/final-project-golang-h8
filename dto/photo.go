package dto

import "time"

// "id", "title", "caption", "photourl", "userid", "createdAt", "updatedAt"
type NewPhotoRequest struct {
	Title    string `json:"title" validate:"required"  example:"Charger Iphone"`
	Caption  string `json:"caption" validate:"required"  example:"Charger Iphone"`
	PhotoUrl string `json:"photoUrl" validate:"required"  example:"Charger Iphone"`
	UserId   int    `json:"userId" validate:"required"  example:"1"`
}

type NewPhotoResponse struct {
	Result     string `json:"result"  example:"Charger Iphone"`
	StatusCode int    `json:"statusCode"  example:"201"`
	Message    string `json:"message"  example:"Charger Iphone"`
}

type PhotoResponse struct {
	Id        int       `json:"id"  example:"1"`
	Title     string    `json:"title"  example:"Charger Iphone"`
	Caption   string    `json:"caption"  example:"Charger Iphone"`
	PhotoUrl  string    `json:"photoUrl"  example:"Charger Iphone"`
	UserId    int       `json:"userId"  example:"1"`
	CreatedAt time.Time `json:"createdAt"  example:"2023-01-01"`
	UpdatedAt time.Time `json:"updatedAt"  example:"2023-01-01"`
}

type PhotoResultResponse struct {
	Result     string          `json:"result"  example:"Charger Iphone"`
	StatusCode int             `json:"statusCode"  example:"201"`
	Message    string          `json:"message"  example:"Charger Iphone"`
	Data       []PhotoResponse `json:"data"`
}
