package main

import router "pgsrv/app/router"

func main() {
	//初始化路由
	router.InitRouter()
	dao.InitMongo()
}
