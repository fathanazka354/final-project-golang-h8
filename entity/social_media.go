package entity

import (
	"final-project/dto"
	"time"
)

type SocialMedia struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"socialMediaUrl"`
	UserId         int       `json:"userId"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (s *SocialMedia) SocialMediaUrlEntityToDto() dto.SocialMediaResponse {
	return dto.SocialMediaResponse{
		Id:             s.Id,
		Name:           s.Name,
		SocialMediaUrl: s.SocialMediaUrl,
		UserId:         s.UserId,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}

}
