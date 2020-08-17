package faq

import (
	"pgsrv/app/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Faq 问题对象
type Faq struct {
	ID     primitive.ObjectID `json:"id" bson:"_id"`
	Title  string             `json:"title" bson:"title"`
	OpenID string             `json:"openId" bson:"openId"`
}

func db() *mongo.Collection {
	return dao.Conn().Collection("faq")
}

func save(faq *Faq) {
	db().InsertOne(nil, bson.M{})
}
