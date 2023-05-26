package application

import (
	"net/http"
	"strconv"

	"github.com/farwater-create/backend/httputils"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GET(ctx *gin.Context) {
	applicationIdStr, exists := ctx.Params.Get("id")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, httputils.NotFoundError)
		return
	}
	applicationId, err := strconv.Atoi(applicationIdStr)
	if err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
		return
	}
	application := &models.Application{
		Model: gorm.Model{
			ID: uint(applicationId),
		},
	}
	db := ctx.MustGet("db").(*gorm.DB)
	tx := db.First(application, applicationId)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, httputils.NotFoundError)
			return
		}
		logrus.Error(tx.Error)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
	}
	ctx.AbortWithStatusJSON(http.StatusOK, application)
}
