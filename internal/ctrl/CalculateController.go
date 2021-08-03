package ctrl

import (
	"SecondProject/internal/handler"
	"SecondProject/internal/model"
	"github.com/gin-gonic/gin"
)

func Calculator(ctx *gin.Context) {
	calculatoStr, _ := ctx.GetPostForm("calculatoStr")
	if calculatoStr == "" {
		ctx.JSON(200, model.Result{Code: "204", Msg: "请输入表达式", Data: nil})
	}
	result := handler.CalculatoStr(calculatoStr)
	ctx.JSON(200, result)
}
