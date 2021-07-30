package handler

import (
	"SecondProject/internal/model"
	"SecondProject/internal/service"
	"strconv"
	"strings"
)

func CalculatoStr(calculatoStr string) (int, error) {

	calculatoStr = strings.Replace(calculatoStr, " ", "", -1)

	//先验证
	_, err := service.ValStrService(calculatoStr)
	if err != nil {
		return 0, err
	} else {
		//进行计算
		result, e := calculate(calculatoStr)
		if e != nil {
			return 0, e
		} else {
			return result, nil
		}
	}
}

func calculate(exp string) (int, error) {
	// 数栈
	numStack := &model.Stack{
		MaxTop: 20,
		Top:    -1,
	}
	// 符号栈
	operStack := &model.Stack{
		MaxTop: 20,
		Top:    -1,
	}

	//exp := "30+30*6-4-6+4"
	// 定义一个 index ，帮助扫描 exp
	index := 0
	// 为了配合运算，需要定义的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""
	for {
		//处理多位数的问题
		ch := exp[index : index+1] // 字符串.
		//ch ==>"+" ===> 43
		temp := int([]byte(ch)[0])  // 就是字符对应的ASCiI码
		if operStack.IsOper(temp) { // 说明是符号
			// 如果 operStack 是一个空栈，直接入栈
			if operStack.Top == -1 { // 空栈
				operStack.Push(temp)
			} else {
				// 如果发现 opertStack 栈顶的运算符的优先级大于等于当前准备入栈的运算符的优先级，
				// 就从符号栈 pop 出，并从数栈也 pop 两个数，进行运算，运算后的结果再重新入栈到数栈，
				// 当前符号再入符号栈
				if operStack.Priority(operStack.Arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					// 将计算结果重新入数栈
					numStack.Push(result)
					// 当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { // 说明是数
			// 处理多位数的思路
			// 1 定义一个变量 keepNum string, 做拼接
			keepNum += ch
			// 2 每次要向 index 的后面字符测试一下，看看是不是运算符，然后处理
			// 如果已经到表达最后，直接将 keepNum 入栈
			if index == len(exp)-1 {

				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				// 向 index 后面测试看看是不是运算符
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}
		// 继续扫描
		// 先判断 index 是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}

	//如果栈顶是*或\则需要再次取值运算
	if operStack.Arr[operStack.Top] == 42 || operStack.Arr[operStack.Top] == 47 {
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		// 将计算结果重新入数栈
		numStack.Push(result)
	}

	unoperStack := operStack.UnStack()
	unnumStack := numStack.UnStack()

	// 如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数。运算后的结果入数栈，直到符号栈为空。
	for {
		if unoperStack.Top == -1 {
			break // 退出条件
		}
		num1, _ = unnumStack.Pop()
		num2, _ = unnumStack.Pop()
		oper, _ = unoperStack.Pop()
		result = unoperStack.Cal(num1, num2, oper)
		// 将计算结果重新入数栈
		unnumStack.Push(result)
	}
	// 结果是 numStack 最后数
	res, _ := unnumStack.Pop()
	return res, nil
}
