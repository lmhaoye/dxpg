package dao

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func InitMongo() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetMaxPoolSize(20)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	// client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// 指定获取要操作的数据集
	db = client.Database("pgdb")
}

/**
* mongo连接对象
**/
func Conn() *mongo.Database {
	return db
}
