package dao

import (
	"context"
	"log"
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
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

// ToBson ...
func ToBson(r interface{}) bson.M {
	result := make(bson.M)
	v := reflect.ValueOf(r)
	t := reflect.TypeOf(r)

	for i := 0; i < v.NumField(); i++ {
		filed := v.Field(i)
		tag := t.Field(i).Tag
		key := tag.Get("bson")
		if key == "" || key == "-" || key == "_id" {
			continue
		}
		keys := strings.Split(key, ",")
		if len(keys) > 0 {
			key = keys[0]
		}
		// TODO: 处理字段嵌套问题
		switch filed.Kind() {
		case reflect.Int, reflect.Int64:
			v := filed.Int()
			if v != 0 {
				result[key] = v
			}
		case reflect.String:
			v := filed.String()
			if v != "" {
				result[key] = v
			}
		case reflect.Bool:
			result[key] = filed.Bool()
		case reflect.Ptr:

		case reflect.Float64:
			v := filed.Float()
			if v != 0 {
				result[key] = v
			}
		case reflect.Float32:
			v := filed.Float()
			if v != 0 {
				result[key] = v
			}
		default:
		}
	}
	return result
}
