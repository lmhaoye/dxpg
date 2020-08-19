package building

import (
	"context"
	"log"
	"pgsrv/app/dao"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Building 对象
type Building struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func db() *mongo.Collection {
	return dao.Conn().Collection("building")
}

func findAll() []*Building {
	var results []*Building
	ctx := context.Background()
	cur, err := db().Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result Building
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return results
}

// getByID 查询单个id
func GetByID(id string) *Building {
	var result Building
	objID, _ := primitive.ObjectIDFromHex(id)
	err := db().FindOne(nil, bson.M{"_id": objID}).Decode(&result)
	if err != nil {
		return nil
	}
	return &result
}
