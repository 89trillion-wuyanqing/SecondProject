package http

import (
	"SecondProject/internal/router"
	"SecondProject/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 初始化函数
func Init() {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()
	// 作为Server的构造器
	engine := gin.Default()
	//var p = &engine
	router.Route(engine)

	httpPort := utils.GetVal("server", "HttpPort")
	err := engine.Run(":" + httpPort)
	if err != nil {
		log.Println("http服务启动失败")
		panic("http服务启动失败")
	}
}
