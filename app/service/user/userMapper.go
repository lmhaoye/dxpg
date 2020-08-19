package user

import (
	"log"
	"pgsrv/app/dao"
	"pgsrv/app/service/building"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User 用户数据
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
	Building  *building.Building `json:"building" bson:"building"`
}

func db() *mongo.Collection {
	return dao.Conn().Collection("user")
}

func saveUser(user *User) {
	openid := user.OpenID
	var dbUserTemp User
	log.Println(bson.M{"openid": openid})
	err := db().FindOneAndUpdate(nil, bson.M{"openid": openid}, bson.M{"$set": bson.M{
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
			res, err := db().InsertOne(nil, bson.M{"openid": openid})
			if err != nil {
				log.Fatalf("save user error=%v", err)
			}
			id := res.InsertedID
			log.Printf("save user success id:%v", id)
			return
		}
		log.Println(err)
	}
	return
}

func getByOpenid(openid string) *User {
	var user User
	err := db().FindOne(nil, bson.M{"openid": openid}).Decode(&user)
	if err != nil {
		log.Printf("decode error=%v", err)
		return nil
	}
	return &user
}

func saveUserBuilding(openid string, buildingID string) {
	bd := building.GetByID(buildingID)
	_, err := db().UpdateOne(nil, bson.M{"openid": openid}, bson.M{
		"building": *bd,
	})
	if err != nil {
		log.Fatalf("save building error =%v", err)
	}
}
