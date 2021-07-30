package router

import (
	"SecondProject/internal/ctrl"
	"github.com/gin-gonic/gin"
)

func Route(engine *gin.Engine) {
	engine.POST("/calculator", ctrl.Calculator)
}
