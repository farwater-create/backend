package main

import (
	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/middleware"
	"github.com/farwater-create/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use()
	r.POST("/user", middleware.ApiKeyMiddleware([]string{apiperms.PostUser}), routes.PostUser)
	r.GET("/user/:id", middleware.ApiKeyMiddleware([]string{apiperms.GetUser}), routes.GetUser)
	r.Run()
}
