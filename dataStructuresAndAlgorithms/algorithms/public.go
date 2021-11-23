package algorithms

import (
	"fmt"
)

// 求一个三元表达式的值
func Opt(a int, b int, opt rune) int {
	if opt == '+' {
		return a + b
	} else if opt == '-' {
		return a - b
	} else if opt == '*' {
		return a * b
	} else if opt == '/' {
		return a / b
	} else {
		fmt.Println("ERROR INPUT!!!")
		return INF
	}
}

// 求两个数的最大值
func GetTheMax(a int, b int) int {
	// 注意：GO语言抛弃了三元运算符，以下写法是错误的
	//return a > b ? a : b
	if a > b {
		return a
	} else if a < b {
		return b
	} else {
		return a
	}
}
