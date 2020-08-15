package service

import (
	"encoding/json"
	"log"
	"pgsrv/app/dao"
	"pgsrv/app/define"
	"pgsrv/app/wechat"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	var rsMap map[string]string
	err := json.Unmarshal([]byte(rs), &rsMap)

	if err != nil {
		log.Fatalln("json error=%v", err)
	}
	openid := rsMap["openid"]
	user := &User{
		OpenId: openid,
	}
	saveUser(user)

	c.JSON(200, define.ReturnOk(rsMap))
}
func HandlerUserGet(c *gin.Context) {

	c.JSON(200, define.ReturnDefault("login success"))
}

func HandlerUserUpdate(c *gin.Context) {
	body := &User{}
	c.Bind(body)
	log.Println(body)
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

func dbUser() *mongo.Collection {
	return dao.Conn().Collection("user")
}

func saveUser(user *User) {
	openid := user.OpenId
	dbData := getByOpenid(openid)
	if dbData != nil {
		return
	}

	res, err := dbUser().InsertOne(nil, bson.M{"openid": openid})
	if err != nil {
		log.Fatalln("save user error=%v", err)
	}
	id := res.InsertedID
	log.Println("save user success id:%v", id)

}

func getByOpenid(openid string) *User {
	var user User
	err := dbUser().FindOne(nil, bson.M{"openid": openid}).Decode(&user)
	if err != nil {
		log.Printf("decode error=%v", err)
		return nil
	}
	return &user
}
