package social_media_repository

import (
	"final-project/entity"
	"final-project/pkg/errs"
)

type SocialMediaRepository interface {
	CreateSocialMedia(newProduct entity.SocialMedia) errs.MessageErr
	GetSocialMedias() ([]*entity.SocialMedia, errs.MessageErr)
	GetSocialMediaById(socialMediaId int) (*entity.SocialMedia, errs.MessageErr)
	UpdateSocialMedia(payload entity.SocialMedia) errs.MessageErr
	DeleteSocialMedia(socialMediaId int) errs.MessageErr
}
