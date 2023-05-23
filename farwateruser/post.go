package farwateruser

import (
	"errors"
	"net/http"
	"time"

	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/httputils"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var ErrorUserExists = errors.New("discord id or minecraft uuid already exists")
var ErrorInternal = errors.New("internal server error")

type PostUserInput struct {
	DiscordID     string    `json:"discordID" validate:"required"`
	MinecraftUUID string    `json:"minecraftUUID" validate:"required"`
	Birthday      time.Time `json:"birthday" validate:"required"`
}

func POST(ctx *gin.Context) {
	postUserInput := &PostUserInput{}

	if !httputils.BindJSON(ctx, postUserInput) {
		return
	}

	user := &models.User{
		DiscordID:     postUserInput.DiscordID,
		MinecraftUUID: postUserInput.MinecraftUUID,
		Birthday:      postUserInput.Birthday,
	}
	db := ctx.MustGet("db").(*gorm.DB)
	tx := db.Where("discord_id = (?) OR minecraft_uuid = (?)", postUserInput.DiscordID, postUserInput.MinecraftUUID).First(user)
	exists := tx.RowsAffected > 0

	if exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "discord id or minecraft uuid already exists",
		})
		return
	}

	tx = db.Create(user)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
		return
	}

	TriggerEvent(EventRegister, user)

	permissions, ok := httputils.ApiTokenPermissions(ctx)
	if !ok {
		return
	}

	json := gin.H{
		"id":        user.ID,
		"createdAt": user.CreatedAt,
	}

	if permissions[apiperms.UserAge] {
		json["birthday"] = user.Birthday
	}

	ctx.AbortWithStatusJSON(http.StatusOK, json)
}
