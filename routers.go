package main

import (
	"ginEssential/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Resgister)
	return r
}
