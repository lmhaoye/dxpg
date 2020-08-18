package wechat

import (
	"fmt"
	"log"
	"pgsrv/app/util"
	"pgsrv/conf"
)

func AuthCode2Session(code string) string {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", conf.Wh.AppID, conf.Wh.Secret, code)
	rs := util.Get(url)
	log.Println(url)
	return rs
}
