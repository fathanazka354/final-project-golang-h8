package dto

import "time"

type NewSocialMediaRequest struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"socialMediaUrl" validate:"required"`
	UserId         int    `json:"userId"`
}
type NewSocialMediaResponse struct {
	Result     string `json:"result"  example:"Charger Iphone"`
	StatusCode int    `json:"statusCode"  example:"201"`
	Message    string `json:"message"  example:"Charger Iphone"`
}

type SocialMediaResponse struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserId         int       `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
type SocialMediaResultResponse struct {
	Result     string                `json:"result"  example:"Charger Iphone"`
	StatusCode int                   `json:"statusCode"  example:"201"`
	Message    string                `json:"message"  example:"Charger Iphone"`
	Data       []SocialMediaResponse `json:"data"`
}
type SocialMediaOneResultResponse struct {
	Result     string              `json:"result"  example:"Charger Iphone"`
	StatusCode int                 `json:"statusCode"  example:"201"`
	Message    string              `json:"message"  example:"Charger Iphone"`
	Data       SocialMediaResponse `json:"data"`
}
