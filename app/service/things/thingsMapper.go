package things

import (
	"pgsrv/app/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Things 事项
type Things struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Title string             `json:"title" bson:"title"`
}

func db() *mongo.Collection {
	return dao.Conn().Collection("things")
}

// GetByID 获取单个数据
func GetByID(id string) *Things {
	var result Things
	objID, _ := primitive.ObjectIDFromHex(id)
	err := db().FindOne(nil, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}
