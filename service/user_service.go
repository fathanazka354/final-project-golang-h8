package service

import (
	"final-project/dto"
	"final-project/entity"
	"final-project/pkg/errs"
	"final-project/pkg/helpers"
	"final-project/repository/user_repository"
	"net/http"
)

type userService struct {
	userRepo user_repository.UserRepository
}

type UserService interface {
	CreateNewUser(newUserRequest dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr)
	Login(payload dto.NewUserRequestLogin) (*dto.LoginResponse, errs.MessageErr)
}

func NewUserService(userRepo user_repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Login(payload dto.NewUserRequestLogin) (*dto.LoginResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	user, err := u.userRepo.GetUserByEmail(payload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, errs.NewUnauthenticatedError("invalid email/password")
		}
		return nil, err
	}
	isValidPassword := user.ComparePassword(payload.Password)

	if !isValidPassword {
		return nil, errs.NewUnauthenticatedError("invalid email/password")
	}

	response := dto.LoginResponse{
		Result:     "success",
		Message:    "logged in successfully",
		StatusCode: http.StatusOK,
		Data: dto.TokenResponse{
			Token: user.GenerateToken(),
		},
	}

	return &response, nil
}

func (u *userService) CreateNewUser(payload dto.NewUserRequest) (*dto.NewUserResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	userEntity := entity.User{
		Email:    payload.Email,
		Password: payload.Password,
		Age:      payload.Age,
		Username: payload.Username,
	}

	err = userEntity.HashPassword()

	if err != nil {
		return nil, err
	}

	if payload.Age <= 8 {
		return nil, errs.NewBadRequest("age must be than equals 8")
	} else if payload.Username == "" {
		return nil, errs.NewBadRequest("username must be filled")
	}
	err = u.userRepo.CreateNewUser(userEntity)

	if err != nil {
		return nil, err
	}

	response := dto.NewUserResponse{
		Result:     "success",
		Message:    "user registered successfully",
		StatusCode: http.StatusCreated,
	}

	return &response, nil
}
