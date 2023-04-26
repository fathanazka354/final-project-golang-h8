package helpers

import (
	"final-project/pkg/errs"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	err := validator.New().Struct(payload)
	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	return nil
}
