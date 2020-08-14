package service

import (
	"pgsrv/app/define"

	"github.com/gin-gonic/gin"
)

// 默认处理逻辑
func NoHandler(c *gin.Context) {
	c.JSON(200, define.ReturnDefault("404"))
}
