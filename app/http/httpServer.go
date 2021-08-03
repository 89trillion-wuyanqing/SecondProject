package http

import (
	"SecondProject/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

// 初始化函数
func Init() {
	// 作为Server的构造器
	engine := gin.Default()
	//var p = &engine
	router.Route(engine)

	err := engine.Run(":8000")
	if err != nil {
		log.Fatal("http服务器启动失败")
	}

}
