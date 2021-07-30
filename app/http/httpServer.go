package http

import (
	"SecondProject/internal/router"
	"github.com/gin-gonic/gin"
)

// 初始化函数
func Init() {
	// 作为Server的构造器
	engine := gin.Default()
	//var p = &engine
	router.Route(engine)

	engine.Run(":8000")

}
