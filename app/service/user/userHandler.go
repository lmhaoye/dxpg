package user

import (
	"encoding/json"
	"log"
	"pgsrv/app/define"
	"pgsrv/app/wechat"

	"github.com/gin-gonic/gin"
)

// LoginBody 登录body体
type LoginBody struct {
	Code string `form:"code" binding:"required"`
}

// HandlerUserLogin 登录
func HandlerUserLogin(c *gin.Context) {
	body := &LoginBody{}
	c.Bind(body)
	code := body.Code
	rs := wechat.AuthCode2Session(code)
	log.Println(rs)
	var rsMap map[string]string
	err := json.Unmarshal([]byte(rs), &rsMap)

	if err != nil {
		log.Fatalf("json error=%v", err)
	}
	openid := rsMap["openid"]
	user := &User{
		OpenID: openid,
	}
	saveUser(user)

	c.JSON(200, define.ReturnOk(rsMap))
}

// HandlerUserGet 获取用户信息
func HandlerUserGet(c *gin.Context) {
	openid, _ := c.GetQuery("openid")
	user := getByOpenid(openid)
	log.Println(user)

	c.JSON(200, define.ReturnDefault("login success"))
}

// HandlerUserUpdate 更新用户信息
func HandlerUserUpdate(c *gin.Context) {
	body := &User{}
	c.ShouldBind(body)
	if len(body.OpenID) == 0 {
		c.JSON(200, define.ReturnFail("更新失败，无openid"))
		return
	}
	log.Println(body)
	saveUser(body)
	c.JSON(200, define.ReturnDefault("update success"))
}

// HandlerUserSet 设置楼栋与棚改意见
func HandlerUserSet(c *gin.Context) {
	c.JSON(200, define.ReturnDefault("ok"))
}
