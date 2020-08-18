package main

import (
	"log"
	dao "pgsrv/app/dao"
	router "pgsrv/app/router"
	"pgsrv/conf"
)

func main() {
	conf.InitConfig()
	//初始化路由
	dao.InitMongo()
	router.InitRouter()

}

// CheckError 检查错误
func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}
