package conf

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var Ser = &Server{}
var Wh = &Wechat{}

type Config struct {
	Server Server `mapstructure:"server"`
	Wechat Wechat `mapstructure:"wechat"`
}

type Server struct {
	Port int64 `mapstructure:"port"`
}

type Wechat struct {
	AppID  string `mapstructure:"appid"`
	Secret string `mapstructure:"secret"`
}

// InitConfig 初始化配置
func InitConfig() {
	c := &Config{}
	filename := "./conf/conf.yaml"

	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config is failed err:", err)
	}

	err = viper.Unmarshal(c)
	if err != nil {
		fmt.Println("unmarshal config is failed, err:", err)
	}

	log.Printf("config read success ==> %v", *c)
	Ser = &c.Server
	Wh = &c.Wechat
}
