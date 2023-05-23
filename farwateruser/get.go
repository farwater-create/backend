package farwateruser

import (
	"net/http"

	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/httputils"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GET(ctx *gin.Context) {
	query, exists := ctx.Params.Get("id")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httputils.BadRequestError)
		return
	}

	user := &models.User{}
	db := ctx.MustGet("db").(*gorm.DB)
	tx := db.Where("discord_id = (?) OR minecraft_uuid = (?) OR id = (?)", query, query, query).Find(user)
	if tx.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, httputils.NotFoundError)
		return
	}

	permissions, ok := httputils.ApiTokenPermissions(ctx)

	if !ok {
		return
	}
	json :=
		gin.H{
			"discordID":     user.DiscordID,
			"minecraftUUID": user.MinecraftUUID,
			"id":            user.ID,
		}

	if permissions[apiperms.UserAge] {
		json["birthday"] = user.Birthday
	}

	ctx.AbortWithStatusJSON(http.StatusOK, json)
}
