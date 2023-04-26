package handler

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/helpers"
	"final-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) photoHandler {
	return photoHandler{
		photoService: photoService,
	}
}

// @BasePath /photos
// CreateNewPhoto godoc
// @Summary create photo
// @Schemes
// @Description create a new photo
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} CreatePhoto
// @Router /photos [post]
func (ph photoHandler) CreateNewPhoto(ctx *gin.Context) {
	var newPhotoRequest dto.NewPhotoRequest

	if err := ctx.ShouldBindJSON(&newPhotoRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	user := ctx.MustGet("userData").(entity.User)

	response, err := ph.photoService.CreateNewPhoto(user.Id, newPhotoRequest)
	if err != nil {
		ctx.JSON(err.Status(), gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (ph *photoHandler) GetPhotoById(ctx *gin.Context) {
	photoId, err := helpers.GetParamById(ctx, "photoId")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	response, err := ph.photoService.GetPhotoById(photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (ph *photoHandler) UpdatePhoto(ctx *gin.Context) {
	photoId, err := helpers.GetParamById(ctx, "photoId")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	var newPhotoRequest dto.NewPhotoRequest
	if err := ctx.ShouldBindJSON(&newPhotoRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	response, err := ph.photoService.UpdatePhoto(photoId, newPhotoRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (ph *photoHandler) GetPhotos(ctx *gin.Context) {
	response, err := ph.photoService.GetPhotos()
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (ph *photoHandler) DeletePhoto(ctx *gin.Context) {
	photoId, err := helpers.GetParamById(ctx, "photoId")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	response, err := ph.photoService.DeletePhoto(photoId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(response.StatusCode, response)
}
