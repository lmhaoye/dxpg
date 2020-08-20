package things

import (
	"pgsrv/app/define"

	"github.com/gin-gonic/gin"
)

// HandlerThingsGet 获取单个对象
func HandlerThingsGet(c *gin.Context) {
	c.JSON(200, define.ReturnDefault("ok"))
}
