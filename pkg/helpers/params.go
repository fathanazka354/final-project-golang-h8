package helpers

import (
	"final-project/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamById(c *gin.Context, key string) (int, errs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, errs.NewUnprocessibleEntityError("invalid id product")
	}

	return id, nil
}
