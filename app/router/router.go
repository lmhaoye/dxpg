package router

import (
	"pgsrv/app/service"
	"pgsrv/app/service/building"
	"pgsrv/app/service/faq"
	"pgsrv/app/service/user"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由
func InitRouter() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.NoRoute(service.NoHandler)
	userGroup := router.Group("user")
	{
		userGroup.POST("/login", user.HandlerUserLogin)
		userGroup.POST("/update", user.HandlerUserUpdate)
		userGroup.GET("/get", user.HandlerUserGet)
	}
	branch := router.Group("building")
	{
		branch.GET("/all", building.HandlerGetAll)
	}
	faqGroup := router.Group("faq")
	{
		faqGroup.GET("/list", faq.HandlerPageFAQ)
		faqGroup.POST("/submit", faq.HandlerSubmitFAQ)
	}
	router.Run()
}
