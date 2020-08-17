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
		log.Fatalf("json error=%v", err)
	}
	openid := rsMap["openid"]
	user := &User{
		OpenID: openid,
	}
	saveUser(user)

	c.JSON(200, define.ReturnOk(rsMap))
}
func HandlerUserGet(c *gin.Context) {
	openid, _ := c.GetQuery("openid")
	user := getByOpenid(openid)
	log.Println(user)

	c.JSON(200, define.ReturnDefault("login success"))
}

func HandlerUserUpdate(c *gin.Context) {
	body := &User{}
	c.ShouldBind(body)
	log.Println(body)
	saveUser(body)
	c.JSON(200, define.ReturnDefault("login success"))
}

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	OpenID    string             `json:"openId" bson:"openId"`
	NickName  string             `json:"nickName" bson:"nickName"`
	Gender    string             `json:"gender" bson:"gender"`
	City      string             `json:"city" bson:"city"`
	Province  string             `json:"province" bson:"province"`
	Country   string             `json:"country" bson:"country"`
	AvatarURL string             `json:"avatarUrl" bson:"avatarUrl"`
	UnionID   string             `json:"unionId" bson:"unionId"`
}

func dbUser() *mongo.Collection {
	return dao.Conn().Collection("user")
}

func saveUser(user *User) {
	openid := user.OpenID
	var dbUserTemp User
	err := dbUser().FindOneAndUpdate(nil, bson.M{"openid": openid}, bson.M{"$set": bson.M{
		"nickName":  user.NickName,
		"gender":    user.Gender,
		"city":      user.City,
		"province":  user.Province,
		"country":   user.Country,
		"avatarUrl": user.AvatarURL},
	}).Decode(&dbUserTemp)
	if err != nil {
		//如果没有数据
		if err == mongo.ErrNoDocuments {
			res, err := dbUser().InsertOne(nil, bson.M{"openid": openid})
			if err != nil {
				log.Fatalln("save user error=%v", err)
			}
			id := res.InsertedID
			log.Println("save user success id:%v", id)
			return
		}
		log.Println(err)
	}
	return
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
