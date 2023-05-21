package kofi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/farwater-create/backend/models"
	"github.com/gin-gonic/gin"
)

type KofiPaymentRequestForm struct {
	Data string `form:"data"`
}

func POST(ctx *gin.Context) {
	contentType := ctx.GetHeader("Content-Type")
	if contentType != "application/x-www-form-urlencoded" {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	kofiPaymentRequestForm := &KofiPaymentRequestForm{}
	ctx.ShouldBind(kofiPaymentRequestForm)
	kofiShopOrder := &models.KofiShopOrder{}
	err := json.Unmarshal([]byte(kofiPaymentRequestForm.Data), kofiShopOrder)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
	listeners := kofiShopOrderEventListeners[KofiEvent(kofiShopOrder.Type)]
	for _, listener := range listeners {
		listener(kofiShopOrder)
	}
}
