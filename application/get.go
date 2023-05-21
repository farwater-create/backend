package application

import (
	"net/http"

	"github.com/farwater-create/backend/httputils"
	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GET(ctx *gin.Context) {
	applicationId, exists := ctx.Params.Get("id")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusNotFound, httputils.NotFoundError)
		return
	}
	application := &models.Application{}
	tx := models.DB.Find(application, "WHERE id = (?)", applicationId)
	if tx.Error != nil {
		logrus.Error(tx.Error)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httputils.InternalServerError)
	}
	if tx.RowsAffected <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, httputils.NotFoundError)
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, application)
}
