package main

import (
	"log"
	dao "pgsrv/app/dao"
	router "pgsrv/app/router"
)

func main() {
	//初始化路由
	dao.InitMongo()
	router.InitRouter()

}

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}
