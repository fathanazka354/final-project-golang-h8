package service

import (
	"final-project/dto"
	"final-project/pkg/errs"
	"final-project/repository/photo_repository"
)

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

type PhotoService interface {
	CreateNewPhoto(payload dto.NewPhotoRequest) errs.MessageErr
	GetPhotoById(photoId int) (*dto.PhotoResponse, errs.MessageErr)
	GetPhotos() ([]*dto.PhotoResultResponse, errs.MessageErr)
	UpdatePhoto(photoId int, payload dto.NewPhotoRequest) errs.MessageErr
	DeletePhoto(photoId int) errs.MessageErr
}
