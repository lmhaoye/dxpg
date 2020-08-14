package main

import (
	dao "pgsrv/app/dao"
	router "pgsrv/app/router"
)

func main() {
	//初始化路由
	dao.InitMongo()
	router.InitRouter()

}
