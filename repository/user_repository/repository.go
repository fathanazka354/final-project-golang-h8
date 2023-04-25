package user_repository

import (
	"final-project/entity"
	"final-project/pkg/errs"
)

type UserRepository interface {
	CreateNewUser(payload entity.User) errs.MessageErr
	GetUserById(userId int) (*entity.User, errs.MessageErr)
	GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr)
}
