package wechat

import (
	"fmt"
	"log"
	"pgsrv/app/define"
	"pgsrv/app/util"
)

func AuthCode2Session(code string) string {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%v&secret=%v&js_code=%v&grant_type=authorization_code", define.APPID, define.SECRET, code)
	rs := util.Get(url)
	log.Println(url)
	return rs
}
