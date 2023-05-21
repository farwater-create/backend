package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
)

var validate = validator.New()

/*
example

	if !models.BindJSON(ctx, postUserInput) {
		return
	}
*/
func BindJSON(ctx *gin.Context, inputModel any) bool {
	err := ctx.BindJSON(inputModel)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BadRequestError)
		logrus.Info(err)
		return false
	}
	if err = validate.Struct(inputModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BadRequestError)
		logrus.Info(err)
		return false
	}
	return true
}
