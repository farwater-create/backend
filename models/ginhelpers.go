package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var InternalServerError = gin.H{
	"error": "internal server error",
}

var ConflictError = gin.H{
	"error": "entry exists",
}

var BadRequestError = gin.H{
	"error": "bad request",
}

/*
example

	var modelInput = &modelInput{}
	if !models.BindJSON(ctx, modelInput) {
		return
	}
	var model = &model{
		...model input fields
	}
	Create(ctx, model)
*/
func Create(ctx *gin.Context, model any) {
	tx := DB.Create(model)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, model)
}

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
	}
	if err = validate.Struct(inputModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BadRequestError)
		return false
	}
	return true
}
