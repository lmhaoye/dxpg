package router

import (
	service "pgsrv/app/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.NoRoute(service.NoHandler)
	userGroup := router.Group("user")
	{
		userGroup.POST("/login", service.HandlerUserLogin)
		userGroup.POST("/update", service.HandlerUserUpdate)
		userGroup.GET("/get", service.HandlerUserGet)
	}
	branch := router.Group("building")
	{
		branch.GET("/all", service.HandlerGetAll)
	}
	router.Run()
}
