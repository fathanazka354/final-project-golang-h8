package handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/helpers"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) socialMediaHandler {
	return socialMediaHandler{
		socialMediaService: socialMediaService,
	}
}
func (socialMediaHandler *socialMediaHandler) CreateSocialMedia(ctx *gin.Context) {
	var newSocialMediaRequest dto.NewSocialMediaRequest

	if err := ctx.ShouldBindJSON(&newSocialMediaRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := socialMediaHandler.socialMediaService.CreateSocialMedia(newSocialMediaRequest.UserId, newSocialMediaRequest)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)

}
func (socialMediaHandler *socialMediaHandler) GetSocialMedias(ctx *gin.Context) {

	response, err := socialMediaHandler.socialMediaService.GetSocialMedias()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)

}
func (socialMediaHandler *socialMediaHandler) GetSocialMediaById(ctx *gin.Context) {
	socialMediaId, err := helpers.GetParamById(ctx, "socialMediaId")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	response, err := socialMediaHandler.socialMediaService.GetSocialMediaById(socialMediaId)
	if err != nil {
		ctx.AbortWithStatusJSON(response.StatusCode, err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
func (socialMediaHandler *socialMediaHandler) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaId, err := helpers.GetParamById(ctx, "socialMediaId")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	var socialMediaRequest dto.NewSocialMediaRequest
	if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	user := ctx.MustGet("userData").(entity.User)
	response, err := socialMediaHandler.socialMediaService.UpdateSocialMedia(user.Id, socialMediaId, socialMediaRequest)
	if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
		ctx.JSON(response.StatusCode, response)
		return
	}

	ctx.JSON(response.StatusCode, response)

}
func (socialMediaHandler *socialMediaHandler) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, err := helpers.GetParamById(ctx, "socialMediaId")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	var socialMediaRequest dto.NewSocialMediaRequest
	if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	response, err := socialMediaHandler.socialMediaService.DeleteSocialMedia(socialMediaId)
	if err := ctx.ShouldBindJSON(&socialMediaRequest); err != nil {
		ctx.JSON(response.StatusCode, response)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
