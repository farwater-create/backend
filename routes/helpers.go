package routes

import (
	"net/http"

	"github.com/farwater-create/backend/apiperms"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
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
	}
	if err = validate.Struct(inputModel); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BadRequestError)
		return false
	}
	return true
}

func ApiTokenPermissions(ctx *gin.Context) (permissions map[apiperms.ApiPermission]bool, ok bool) {
	_p, exists := ctx.Get("permissions")
	permissions, ok = _p.(map[apiperms.ApiPermission]bool)
	if !exists || !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError)
		return permissions, false
	}
	return permissions, true
}
