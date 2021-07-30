package service

import (
	"errors"
)

func ValStrService(valStr string) (bool, error) {
	//验证字符串是否规范，不允许出现0-9 + - * /这些字符之外的字符
	//不允许+ - * /连续
	//开头和结尾不允许使用操作符
	// 符号栈
	lastStr := -1  //上一位字符asscii
	indexStr := -1 //当前字符asscii
	index := 0
	for {
		//处理多位数的问题
		ch := valStr[index : index+1] // 字符串.
		//ch ==>"+" ===> 43
		temp := int([]byte(ch)[0]) // 就是字符对应的ASCiI码
		flag := ValLegalStr(temp)
		if !flag {
			//非法字符不允许
			return false, errors.New("存在非法字符，请重新输入！")
		} else {

			//字符合法
			//判断开头和结尾不允许是操作符
			if index == 0 || index+1 == len(valStr) {
				isOperation := ValOperation(temp)
				if isOperation {
					return false, errors.New("开头和结尾不允许使用操作符，不符合规范")
				}
			}

			if indexStr == -1 {
				indexStr = temp
			} else {
				lastStr = indexStr
				indexStr = temp
			}
			isOperation := ValOperation(temp)
			lastOperation := ValOperation(lastStr)
			if isOperation && lastOperation {
				return false, errors.New("操作符连续，不符合规范！")
			}
		}
		// 继续扫描
		// 先判断 index 是否已经扫描到计算表达式的最后
		if index+1 == len(valStr) {
			break
		}
		index++
	}
	return true, nil

}

func ValLegalStr(asscii int) bool {
	if (asscii >= 48 && asscii <= 57) || (asscii == 42 || asscii == 43 || asscii == 45 || asscii == 47) {
		return true
	} else {
		return false
	}
}

func ValOperation(asscii int) bool {
	if asscii == 42 || asscii == 43 || asscii == 45 || asscii == 47 {
		return true
	} else {
		return false
	}
}

//去除空格
