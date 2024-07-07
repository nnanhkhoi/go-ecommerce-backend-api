package initialize

import (
	"fmt"

	"github.com/nnanhkhoi/go-ecommerce-backend-api/global"
)

func Run() {
	// load config
	loadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading config from mysql", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()
	r := InitRouter()

	r.Run(":8002")
}
