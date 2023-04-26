package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helpers"
	"final-project/repository/photo_repository"
	"fmt"
	"net/http"
)

type photoService struct {
	photoRepo photo_repository.PhotoRepository
}

type PhotoService interface {
	CreateNewPhoto(userId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	GetPhotoById(photoId int) (*dto.PhotoResponse, errs.MessageErr)
	GetPhotos() (*dto.PhotoResultResponse, errs.MessageErr)
	UpdatePhoto(photoId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr)
	DeletePhoto(photoId int) (*dto.NewPhotoResponse, errs.MessageErr)
}

func NewPhotoService(photoRepo photo_repository.PhotoRepository) PhotoService {
	return &photoService{
		photoRepo: photoRepo,
	}
}

func (photoService *photoService) CreateNewPhoto(userId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	photoRequest := &entity.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserId:   userId,
	}

	_, err = photoService.photoRepo.CreateNewPhoto(photoRequest)

	if err != nil {
		return nil, err
	}
	response := dto.NewPhotoResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "new photo successfully created",
	}

	return &response, nil
}

func (photoService *photoService) GetPhotoById(photoId int) (*dto.PhotoResponse, errs.MessageErr) {
	response, err := photoService.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, err
	}

	photoResponse := dto.PhotoResponse{
		Id:        response.Id,
		Title:     response.Title,
		Caption:   response.Caption,
		PhotoUrl:  response.PhotoUrl,
		UserId:    response.UserId,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}
	return &photoResponse, nil
}

func (photoService *photoService) GetPhotos() (*dto.PhotoResultResponse, errs.MessageErr) {
	response, err := photoService.photoRepo.GetPhotos()

	fmt.Println("here1")
	if err != nil {
		return nil, err
	}
	fmt.Println("here2")

	var photoResponses []dto.PhotoResponse

	for _, eachPhoto := range response {
		photoResponses = append(photoResponses, eachPhoto.PhotoEntityToDto())
	}
	photoResultResponse := dto.PhotoResultResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "data photos successfully to send",
		Data:       photoResponses,
	}
	return &photoResultResponse, nil
}

func (photoService *photoService) UpdatePhoto(photoId int, payload dto.NewPhotoRequest) (*dto.NewPhotoResponse, errs.MessageErr) {

	_, err := photoService.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, errs.NewNotFoundError("no product exists")
	}
	photo := &entity.Photo{
		Title:    payload.Title,
		Caption:  payload.Caption,
		PhotoUrl: payload.PhotoUrl,
		UserId:   payload.UserId,
	}
	err = photoService.photoRepo.UpdatePhoto(photoId, *photo)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response := dto.NewPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "data photos successfully to update",
	}

	return &response, nil
}

func (photoService *photoService) DeletePhoto(photoId int) (*dto.NewPhotoResponse, errs.MessageErr) {
	_, err := photoService.photoRepo.GetPhotoById(photoId)

	if err != nil {
		return nil, errs.NewNotFoundError("no product exists")
	}

	err = photoService.photoRepo.DeletePhoto(photoId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response := dto.NewPhotoResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "data photos successfully to delete",
	}

	return &response, nil
}
