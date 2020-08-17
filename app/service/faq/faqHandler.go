package faq

import (
	"pgsrv/app/define"

	"github.com/gin-gonic/gin"
)

// HandlerPageFAQ 问题分页
func HandlerPageFAQ(c *gin.Context) {
	c.JSON(200, define.ReturnOk(nil))
}

// HandlerSubmitFAQ 提交问题
func HandlerSubmitFAQ(c *gin.Context) {
	c.JSON(200, define.ReturnOk(nil))
}
