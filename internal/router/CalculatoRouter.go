package router

import (
	"SecondProject/internal/ctrl"
	"github.com/gin-gonic/gin"
)

//路由层
func Route(engine *gin.Engine) {
	engine.POST("/calculator", ctrl.Calculator)
}
