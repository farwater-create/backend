package main

import (
	"github.com/farwater-create/backend/apiperms"
	"github.com/farwater-create/backend/config"
	"github.com/farwater-create/backend/database"
	"github.com/farwater-create/backend/farwateruser"
	"github.com/farwater-create/backend/kofi"
	"github.com/farwater-create/backend/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	db := database.New(config.MARIADB_DSN)
	r.Use(middleware.DBMiddleware(db))
	r.POST("/v1/user", middleware.ApiKeyMiddleware([]string{apiperms.PostUser}), farwateruser.POST)
	r.GET("/v1/user/:id", middleware.ApiKeyMiddleware([]string{apiperms.GetUser}), farwateruser.GET)
	r.POST("/v1/application", middleware.ApiKeyMiddleware([]string{apiperms.Applications}))
	r.POST("/v1/application/:id", middleware.ApiKeyMiddleware([]string{apiperms.Applications}))
	r.GET("/v1/apikey", middleware.ApiKeyMiddleware([]string{apiperms.Grant}))
	r.POST("/v1/kofi", kofi.POST)
	r.Run()
}
