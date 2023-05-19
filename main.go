package main

import (
	"github.com/farwater-create/backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/user", routes.PostUser)
	r.GET("/user/:id", routes.GetUser)
	r.Run()
}
