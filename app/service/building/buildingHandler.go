package building

import (
	"pgsrv/app/define"

	"github.com/gin-gonic/gin"
)

// HandlerGetAll 获取所有楼层信息
func HandlerGetAll(c *gin.Context) {
	ret := define.ReturnOk(findAll())
	c.JSON(200, ret)
}
