package handler

import (
	"final-project/database"
	"final-project/repository/photo_repository/photo_pg"
	"final-project/repository/user_repository/user_pg"
	"final-project/service"

	"github.com/gin-gonic/gin"
)

const PORT = "4000"

func StartApp() {
	database.InitializeDatabase()
	db := database.GetDatabaseInstance()
	userRepo := user_pg.NewUserPG(db)
	photoRepo := photo_pg.NewPhotoPG(db)
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo, photoRepo)
	userHandler := NewUserHandler(userService)
	route := gin.Default()

	userRoutes := route.Group("/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
	}

	photoRoutes := route.Group("/photos")
	{
		photoRoutes.POST("/", authService.Authentication())
	}
	route.Run(":" + PORT)
}
