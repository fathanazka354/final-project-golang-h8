package handler

import (
	"final-project/database"
	"final-project/docs"
	"final-project/repository/comment_repository/comment_pg"
	"final-project/repository/photo_repository/photo_pg"
	"final-project/repository/social_media_repository/social_media_pg"
	"final-project/repository/user_repository/user_pg"
	"final-project/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const PORT = "4000"

// @title MyGram API
// @version 1.0
// @description This is a sample service for managing final project
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licences/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
func StartApp() {
	database.InitializeDatabase()
	db := database.GetDatabaseInstance()
	userRepo := user_pg.NewUserPG(db)
	photoRepo := photo_pg.NewPhotoPG(db)
	socialMediaRepo := social_media_pg.NewSocialMediaPG(db)
	commentRepo := comment_pg.NewCommentPG(db)
	userService := service.NewUserService(userRepo)
	photoService := service.NewPhotoService(photoRepo)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	commentService := service.NewCommentService(commentRepo)
	authService := service.NewAuthService(userRepo, photoRepo, socialMediaRepo)
	userHandler := NewUserHandler(userService)
	photoHandler := NewPhotoHandler(photoService)
	commentHandler := NewCommentHandler(commentService)
	socialMediaHandler := NewSocialMediaHandler(socialMediaService)
	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Title = "Product API"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:4000"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoutes := route.Group("/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}

	socialMediaRoutes := route.Group("/social-medias")
	{
		socialMediaRoutes.GET("/", authService.Authentication(), socialMediaHandler.GetSocialMedias)
		socialMediaRoutes.GET("/:socialMediaId", authService.Authentication(), socialMediaHandler.GetSocialMediaById)
		socialMediaRoutes.POST("/", authService.Authentication(), socialMediaHandler.CreateSocialMedia)
		socialMediaRoutes.PUT("/:socialMediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialMediaHandler.UpdateSocialMedia)
		socialMediaRoutes.DELETE("/:socialMediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialMediaHandler.DeleteSocialMedia)
	}

	commentRoutes := route.Group("/comment")
	{
		commentRoutes.GET("/", authService.Authentication(), commentHandler.GetComments)
		commentRoutes.GET("/:commentId", authService.Authentication(), commentHandler.GetCommentById)
		commentRoutes.POST("/", authService.Authentication(), commentHandler.CreateComment)
		commentRoutes.PUT("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.UpdateComment)
		commentRoutes.DELETE("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.DeleteComment)
	}
	photoRoutes := route.Group("/photos")
	{
		photoRoutes.POST("/", authService.Authentication(), authService.Authentication(), photoHandler.CreateNewPhoto)
		photoRoutes.GET("/", authService.Authentication(), photoHandler.GetPhotos)
		photoRoutes.GET("/:photoId", authService.Authentication(), photoHandler.GetPhotoById)
		photoRoutes.PUT("/:photoId", authService.Authentication(), authService.Authorization(), photoHandler.UpdatePhoto)
		photoRoutes.DELETE("/:photoId", authService.Authentication(), authService.Authorization(), photoHandler.DeletePhoto)
	}
	route.Run(":" + PORT)
}
