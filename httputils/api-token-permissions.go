package httputils

import (
	"net/http"

	"github.com/farwater-create/backend/apiperms"
	"github.com/gin-gonic/gin"
)

func ApiTokenPermissions(ctx *gin.Context) (permissions map[apiperms.ApiPermission]bool, ok bool) {
	_p, exists := ctx.Get("permissions")
	permissions, ok = _p.(map[apiperms.ApiPermission]bool)
	if !exists || !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError)
		return permissions, false
	}
	return permissions, true
}
