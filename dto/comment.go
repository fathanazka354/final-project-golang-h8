package dto

import (
	"time"
)

type NewCommentRequest struct {
	UserId  int    `json:"userId"  example:"1"`
	PhotoId int    `json:"photoId"  example:"1"`
	Message string `json:"message" validate:"required"  example:"halo apa kabar?"`
}

type NewCommentResponse struct {
	Result     string `json:"result"  example:"Charger Iphone"`
	StatusCode int    `json:"statusCode"  example:"201"`
	Message    string `json:"message"  example:"Charger Iphone"`
}

type CommentResponse struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userId"`
	PhotoId   int       `json:"photoId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommentResultResponse struct {
	Result     string            `json:"result"  example:"Charger Iphone"`
	StatusCode int               `json:"statusCode"  example:"201"`
	Message    string            `json:"message"  example:"Charger Iphone"`
	Data       []CommentResponse `json:"data"`
}

type CommentResultOneResponse struct {
	Result     string          `json:"result"  example:"Charger Iphone"`
	StatusCode int             `json:"statusCode"  example:"201"`
	Message    string          `json:"message"  example:"Charger Iphone"`
	Data       CommentResponse `json:"data"`
}
