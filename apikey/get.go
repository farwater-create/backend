package apikey

import (
	"net/http"
	"strings"

	"github.com/farwater-create/backend/httputils"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetApiKeyInput struct {
	UserID      uint     `json:"userID" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

func Create(user uint, permissions []string) (*models.ApiKey, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	apiKey := &models.ApiKey{
		UserID:      user,
		Key:         u.String(),
		Permissions: strings.Join(permissions, ";"),
	}
	tx := models.DB.Create(apiKey)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return apiKey, nil
}

func GET(ctx *gin.Context) {
	getApiKeyInput := &GetApiKeyInput{}
	if !httputils.BindJSON(ctx, getApiKeyInput) {
		return
	}
	apiKey, err := Create(getApiKeyInput.UserID, getApiKeyInput.Permissions)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, apiKey)
}
