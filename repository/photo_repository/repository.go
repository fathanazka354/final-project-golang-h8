package photo_repository

import (
	"final-project/entity"
	"final-project/pkg/errs"
)

type PhotoRepository interface {
	CreateNewPhoto(payload *entity.Photo) (*entity.Photo, errs.MessageErr)
	GetPhotoById(photoId int) (*entity.Photo, errs.MessageErr)
	GetPhotos() ([]*entity.Photo, errs.MessageErr)
	UpdatePhoto(photoId int, payload entity.Photo) errs.MessageErr
	DeletePhoto(photoId int) errs.MessageErr
}
