package main

import (
	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/farwateruser"
	"github.com/farwater-create/backend/kofi"
	"github.com/farwater-create/backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/v1/user", middleware.ApiKeyMiddleware([]string{apiperms.PostUser}), farwateruser.POST)
	r.GET("/v1/user/:id", middleware.ApiKeyMiddleware([]string{apiperms.GetUser}), farwateruser.GET)
	r.POST("/v1/kofi", kofi.POST)
	r.Run()
}
