package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helpers"
	"final-project/repository/social_media_repository"
	"net/http"
)

type socialMediaService struct {
	socialMediaRepo social_media_repository.SocialMediaRepository
}

type SocialMediaService interface {
	CreateSocialMedia(userId int, newProduct dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr)
	GetSocialMedias() (*dto.SocialMediaResultResponse, errs.MessageErr)
	GetSocialMediaById(socialMediaId int) (*dto.SocialMediaOneResultResponse, errs.MessageErr)
	UpdateSocialMedia(userId int, socialMediaId int, nw dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr)
	DeleteSocialMedia(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr)
}

func NewSocialMediaService(socialMediaRepo social_media_repository.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{socialMediaRepo: socialMediaRepo}
}

func (socialMediaService *socialMediaService) CreateSocialMedia(userId int, newSocialMedia dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(newSocialMedia)

	if err != nil {
		return nil, err
	}

	payload := entity.SocialMedia{
		Name:           newSocialMedia.Name,
		SocialMediaUrl: newSocialMedia.SocialMediaUrl,
		UserId:         userId,
	}

	err = socialMediaService.socialMediaRepo.CreateSocialMedia(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewSocialMediaResponse{
		Result:     "success",
		StatusCode: http.StatusCreated,
		Message:    "social media created",
	}

	return &response, nil
}

func (socialMediaService *socialMediaService) GetSocialMedias() (*dto.SocialMediaResultResponse, errs.MessageErr) {
	result, err := socialMediaService.socialMediaRepo.GetSocialMedias()
	if err != nil {
		return nil, err
	}

	var socialMedias []dto.SocialMediaResponse
	for _, eachSocialMedia := range result {
		socialMedias = append(socialMedias, eachSocialMedia.SocialMediaUrlEntityToDto())
	}

	response := dto.SocialMediaResultResponse{
		Result:     "success",
		Message:    "social media successfully sent",
		StatusCode: http.StatusOK,
		Data:       socialMedias,
	}

	return &response, nil
}

func (socialMediaService *socialMediaService) GetSocialMediaById(socialMediaId int) (*dto.SocialMediaOneResultResponse, errs.MessageErr) {
	socialMedia, err := socialMediaService.socialMediaRepo.GetSocialMediaById(socialMediaId)
	if err != nil {
		if err.Status() == http.StatusNoContent {
			response := dto.SocialMediaOneResultResponse{
				Result:     "success",
				Message:    "social media successfully sent",
				StatusCode: http.StatusOK,
				Data:       socialMedia.SocialMediaUrlEntityToDto(),
			}
			return &response, nil
		}
		return nil, err
	}

	response := dto.SocialMediaOneResultResponse{
		Result:     "success",
		Message:    "social media successfully sent",
		StatusCode: http.StatusOK,
		Data:       socialMedia.SocialMediaUrlEntityToDto(),
	}
	return &response, nil
}

func (socialMediaService *socialMediaService) UpdateSocialMedia(userId int, socialMediaId int, nw dto.NewSocialMediaRequest) (*dto.NewSocialMediaResponse, errs.MessageErr) {
	payload := entity.SocialMedia{
		Id:             socialMediaId,
		Name:           nw.Name,
		SocialMediaUrl: nw.SocialMediaUrl,
		UserId:         userId,
	}
	_, err := socialMediaService.socialMediaRepo.GetSocialMediaById(socialMediaId)

	if err != nil {
		return nil, errs.NewNotFoundError("no social media exists")
	}
	err = socialMediaService.socialMediaRepo.UpdateSocialMedia(payload)
	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	response := dto.NewSocialMediaResponse{
		Result:     "success",
		Message:    "social media successfully sent",
		StatusCode: http.StatusOK,
	}
	return &response, nil
}
func (socialMediaService *socialMediaService) DeleteSocialMedia(socialMediaId int) (*dto.NewSocialMediaResponse, errs.MessageErr) {

	_, err := socialMediaService.socialMediaRepo.GetSocialMediaById(socialMediaId)

	if err != nil {
		return nil, errs.NewNotFoundError("no social media exists")
	}
	err = socialMediaService.socialMediaRepo.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}
	response := dto.NewSocialMediaResponse{
		Result:     "success",
		Message:    "social media successfully delete",
		StatusCode: http.StatusOK,
	}
	return &response, nil
}
