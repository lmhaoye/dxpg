package service

import "github.com/gin-gonic/gin"

// 默认处理逻辑
func NoHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"success": false,
		"error":   "404 no handler",
		"data":    nil,
	})
}
