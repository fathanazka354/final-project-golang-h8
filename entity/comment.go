package entity

import (
	"final-project/dto"
	"time"
)

type Comment struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	PhotoId   int       `json:"photoId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (comment *Comment) CommentEntityToDto() dto.CommentResponse {
	return dto.CommentResponse{
		Id:        comment.Id,
		UserId:    comment.UserId,
		PhotoId:   comment.PhotoId,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
