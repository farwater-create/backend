package application

import (
	"net/http"

	"github.com/farwater-create/backend/httputils"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PostApplicationInput struct {
	UserID uint   `json:"userId"`
	Reason string `json:"reason" validate:"required"`
}

func POST(ctx *gin.Context) {
	applicationInput := &PostApplicationInput{}
	if !httputils.BindJSON(ctx, applicationInput) {
		return
	}
	application := &models.Application{
		UserID: applicationInput.UserID,
		Status: "pending",
		Reason: applicationInput.Reason,
	}
	tx := models.DB.Create(application)
	if tx.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
		logrus.Error(tx.Error)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, application)
}
