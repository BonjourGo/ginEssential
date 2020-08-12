package main

import (
	"ginEssential/controller"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func GetRouter(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.AuthMiddleWare() ,controller.Info)
	return r
}
