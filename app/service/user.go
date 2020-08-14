package service

import (
	"log"
	"pgsrv/app/define"
	"pgsrv/app/wechat"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginBody struct {
	Code string `form:"code" binding:"required"`
}

func HandlerUserLogin(c *gin.Context) {
	body := &LoginBody{}
	c.Bind(body)
	code := body.Code
	rs := wechat.AuthCode2Session(code)
	log.Println(rs)
	c.JSON(200, define.ReturnDefault("login success"))
}
func HandlerUserGet(c *gin.Context) {
	c.JSON(200, define.ReturnDefault("login success"))
}

func HandlerUserUpdate(c *gin.Context) {
	c.JSON(200, define.ReturnDefault("login success"))
}

type User struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	OpenId    string             `json:"openId" bson:"openId"`
	NickName  string             `json:"nickName" bson:"nickName"`
	Gender    string             `json:"gender" bson:"gender"`
	City      string             `json:"city" bson:"city"`
	Province  string             `json:"province" bson:"province"`
	Country   string             `json:"country" bson:"country"`
	AvatarUrl string             `json:"avatarUrl" bson:"avatarUrl"`
	UnionId   string             `json:"unionId" bson:"unionId"`
}
