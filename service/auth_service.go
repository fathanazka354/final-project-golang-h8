package service

import (
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helpers"
	"final-project/repository/photo_repository"
	"final-project/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
}

type authService struct {
	userRepo  user_repository.UserRepository
	photoRepo photo_repository.PhotoRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		photoRepo: photoRepo,
	}
}

func (a *authService) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(entity.User)

		photoId, err := helpers.GetParamById(ctx, "photoId")

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		photo, err := a.photoRepo.GetPhotoById(photoId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserId != user.Id {
			unauthorizeErr := errs.NewUnauthorizedError("you are authorized to modify the photo data")
			ctx.AbortWithStatusJSON(unauthorizeErr.Status(), unauthorizeErr)
			return
		}

		ctx.Next()

	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user entity.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)
		ctx.Next()
	}
}
