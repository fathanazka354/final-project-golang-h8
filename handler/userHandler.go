package handler

import (
	"final-project/dto"
	"final-project/pkg/errs"
	"final-project/service"
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

// LoginUser godoc
// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param RequestBody body dto.NewUserRequestLogin true "request body json"
// @Success 201 {object} dto.NewUserRequestLogin
// @Router /users/login [post]
func (uh *userHandler) Login(ctx *gin.Context) {
	var newUserRequest dto.NewUserRequestLogin

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := uh.userService.Login(newUserRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
