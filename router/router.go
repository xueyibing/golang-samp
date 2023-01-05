package router

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/api/createUser", controllers.CreateUser)
	router.POST("/api/getPhone", controllers.GetPhone)

	return router
}
