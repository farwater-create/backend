package routes

import (
	"net/http"

	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {
	postUserInput := &models.PostUserInput{}

	if !BindJSON(ctx, postUserInput) {
		return
	}

	user := &models.User{
		DiscordID:     postUserInput.DiscordID,
		MinecraftUUID: postUserInput.MinecraftUUID,
		Birthday:      postUserInput.Birthday,
	}

	tx := models.DB.Where("discord_id = (?) OR minecraft_uuid = (?)", postUserInput.DiscordID, postUserInput.MinecraftUUID).First(user)
	exists := tx.RowsAffected > 0

	if exists {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error": "discord id or minecraft uuid already exists",
		})
		return
	}

	tx = models.DB.Create(user)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, InternalServerError)
		return
	}

	permissions, ok := ApiTokenPermissions(ctx)
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

func GetUser(ctx *gin.Context) {
	query, exists := ctx.Params.Get("id")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, BadRequestError)
		return
	}

	user := &models.User{}

	tx := models.DB.Where("discord_id = (?) OR minecraft_uuid = (?) OR id = (?)", query, query, query).Find(user)
	if tx.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, NotFoundError)
		return
	}

	permissions, ok := ApiTokenPermissions(ctx)

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
