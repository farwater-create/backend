package apikey

import (
	"net/http"

	"github.com/farwater-create/backend/httputils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type GetApiKeyInput struct {
	UserID      uint     `json:"userID" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

func GET(ctx *gin.Context) {
	getApiKeyInput := &GetApiKeyInput{}
	if !httputils.BindJSON(ctx, getApiKeyInput) {
		return
	}
	apiKey, err := New(getApiKeyInput.UserID, getApiKeyInput.Permissions)
	db := ctx.MustGet("db").(*gorm.DB)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
		logrus.Error(err)
		return
	}
	tx := db.Create(apiKey)
	if err = tx.Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
		logrus.Error(err)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, apiKey)
}
