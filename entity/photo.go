package entity

import "time"

type Photo struct {
	Id        int       `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl" validate:"required"`
	UserId    int       `json:"userId" validate:"required,numeric"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
