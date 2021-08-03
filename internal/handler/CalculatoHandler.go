package handler

import (
	"SecondProject/internal/model"
	"SecondProject/internal/service"
	"strings"
)

/**
对表达式进行验证后再计算
*/
func CalculatoStr(calculatoStr string) model.Result {

	calculatoStr = strings.Replace(calculatoStr, " ", "", -1)
	var valResult model.Result
	//先验证
	result := service.ValStrService(calculatoStr)
	if result == valResult {

		r := service.Calculate(calculatoStr)
		return r
	}
	return result

}
