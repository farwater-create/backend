package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiTokenPermissions(ctx *gin.Context) (permissions map[string]bool, ok bool) {
	_p, exists := ctx.Get("permissions")
	permissions, ok = _p.(map[string]bool)
	if !exists || !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError)
		return permissions, false
	}
	return permissions, true
}
