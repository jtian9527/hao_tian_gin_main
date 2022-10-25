
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"haotian_main/config"
	"haotian_main/routers"
	"haotian_main/utils"
)


const (
	address = "localhost:9091"
)


func main() {
	conf, err := config.ParseConfig("./config/config.json")
	if err != nil {
		panic("读取配置文件失败，" + err.Error())
	}
	fmt.Printf("conf:%#v\n", conf)
	utils.InitUcache(conf.RedisConfig)
	engine := gin.Default()
	routers.RegisterRouter(engine)
	engine.Run(":8011") // listen and serve on 0.0.0.0:8080
}
