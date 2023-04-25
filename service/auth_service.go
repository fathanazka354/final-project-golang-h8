package service

import (
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

// func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository) {
// 	return &authService{
// 		userRepo:  userRepo,
// 		photoRepo: photoRepo,
// 	}
// }
