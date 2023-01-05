package router

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/getUserinfo", controllers.CreateUser)

	return router
}
