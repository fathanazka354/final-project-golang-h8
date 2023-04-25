package dto

import "time"

// "id", "title", "caption", "photourl", "userid", "createdAt", "updatedAt"
type NewPhotoRequest struct {
	Title    string `json:"title" validate:"required"`
	Caption  string `json:"captopm" validate:"required"`
	PhotoUrl string `json:"photoUrl" validate:"required"`
	UserId   int    `json:"userId" validate:"required"`
}

type NewPhotoResponse struct {
	Result     string `json:"result"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type PhotoResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"captopm"`
	PhotoUrl  string    `json:"photoUrl"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PhotoResultResponse struct {
	Result     string          `json:"result"`
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       []PhotoResponse `json:"data"`
}
