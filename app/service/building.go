package service

import (
	"context"
	"log"
	"pgsrv/app/dao"
	"pgsrv/app/define"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlerGetAll(c *gin.Context) {
	ret := define.ReturnOk(dbSelectAll())
	c.JSON(200, ret)
}

type Building struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

func dbBuilding() *mongo.Collection {
	return dao.Conn().Collection("building")
}

func dbSelectAll() []*Building {
	var results []*Building
	ctx := context.Background()
	cur, err := dbBuilding().Find(ctx, bson.D{})
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
