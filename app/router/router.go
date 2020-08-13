package router

import (
	service "pgsrv/app/service"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.NoRoute(service.NoHandler)
	userGroup := router.Group("user")
	{
		userGroup.POST("/update", service.UpdateUser)
		userGroup.GET("/get", service.GetUser)
	}
	branch := router.Group("building")
	{
		branch.GET("/all", service.HandlerGetAll)
	}
	router.Run()
}
