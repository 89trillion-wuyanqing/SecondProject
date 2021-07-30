package model

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈
type Stack struct {
	MaxTop int     // 栈最大可以存放数个数
	Top    int     // 栈顶, 因为栈顶固定，因此我们直接使用 Top
	Arr    [20]int // 数组模拟栈
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	// 先判断栈是否满
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	// 放入数据
	this.Arr[this.Top] = val
	return
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断栈是否空
	if this.Top == -1 {
		fmt.Println("stack empty!")
		return 0, errors.New("stack empty")
	}

	// 先取值，再 this.Top--
	val = this.Arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈，从栈顶开始遍历
func (this *Stack) List() {
	// 先判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.Arr[i])
	}

}

// 判断一个字符是不是一个运算符[+, - , * , /]
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// [+, - , * , /] 运算
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num1 + num2
	case 45:
		res = num1 - num2
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误.")
	}
	return res
}

// 返回某个运算符的优先级，由程序员定义
// * / 定义为 1
// + - 定义为 0
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func (this *Stack) UnStack() *Stack {
	unstack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	for {
		if this.Top == -1 {
			break // 退出条件
		}
		num, _ := this.Pop()
		unstack.Push(num)
	}
	return unstack
}
