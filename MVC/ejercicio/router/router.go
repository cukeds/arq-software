package router

import (
	"git/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user", controller.GetUsers)
	return r
}
