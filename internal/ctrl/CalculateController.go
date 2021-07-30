package ctrl

import (
	"SecondProject/internal/handler"
	"github.com/gin-gonic/gin"
)

func Calculator(ctx *gin.Context) {
	calculatoStr, _ := ctx.GetPostForm("calculatoStr")
	result, err := handler.CalculatoStr(calculatoStr)
	if err != nil {
		ctx.JSON(200, map[string]interface{}{
			"code":    201,
			"message": "ERROR",
			"data":    err.Error(),
		})
	} else {
		ctx.JSON(200, map[string]interface{}{
			"code":    200,
			"message": "OK",
			"data":    result,
		})
	}
}
