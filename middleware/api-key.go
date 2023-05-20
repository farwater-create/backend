package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware(permissions []string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer")

		if len(authHeader) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		authorization := authHeader[1]
		authorization = strings.TrimSpace(authorization)

		if authorization == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		apiKey := &models.ApiKey{}
		tx := models.DB.Where("`key` = ?", authorization).First(apiKey)
		if tx.RowsAffected <= 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}
		fmt.Println(apiKey.Permissions)
		userPermissions := strings.Split(apiKey.Permissions, ";")
		permissionsMap := make(map[apiperms.ApiPermission]bool)
		for _, p := range userPermissions {
			fmt.Println(p)
			permissionsMap[p] = true
		}
		for _, p := range permissions {
			if _, ok := permissionsMap[p]; !ok {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": "unauthorized",
				})
				return
			}
		}
		ctx.Set("permissions", permissionsMap)
		ctx.Set("user", apiKey.UserID)
		ctx.Next()
	}
}
