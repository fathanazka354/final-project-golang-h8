package handler

import (
	"final-project/dto"
	"final-project/pkg/errs"
	"final-project/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{
		userService: userService,
	}
}

func (uh userHandler) Register(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequest

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response, err := uh.userService.CreateNewUser(newUserRequest)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Message(),
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (uh *userHandler) Login(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequestLogin

	fmt.Println("here1")
	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	fmt.Println("here2")

	result, err := uh.userService.Login(newUserRequest)
	fmt.Println("here3")

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	fmt.Println("here4")
	ctx.JSON(result.StatusCode, result)
}
