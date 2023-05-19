package routes

import (
	"net/http"

	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
)

func PostUser(ctx *gin.Context) {
	postUserInput := &models.PostUserInput{}

	if !models.BindJSON(ctx, postUserInput) {
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

	models.Create(ctx, user)
}

func GetUser(ctx *gin.Context) {
	query, exists := ctx.Params.Get("id")

	if !exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
		})
		return
	}

	user := &models.User{}

	tx := models.DB.Where("discord_id = (?) OR minecraft_uuid = (?) OR id = (?)", query, query, query).Find(user)
	if tx.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"discordID":     user.DiscordID,
		"minecraftUUID": user.MinecraftUUID,
		"id":            user.ID,
	})
}
