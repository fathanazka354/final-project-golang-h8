package entity

import (
	"final-project/dto"
	"time"
)

type Photo struct {
	Id        int       `json:"id" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl" validate:"required"`
	UserId    int       `json:"userId" validate:"required,numeric"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (p *Photo) PhotoEntityToDto() dto.PhotoResponse {
	return dto.PhotoResponse{
		Id:        p.Id,
		Title:     p.Title,
		Caption:   p.Caption,
		PhotoUrl:  p.PhotoUrl,
		UserId:    p.UserId,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
