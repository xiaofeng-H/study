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

// 注意：GO语言抛弃了三元运算符，以下写法是错误的
//return a > b ? a : b
// 求两个数（int）的较大者
func MaxIntAB(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 求两个数（int）的较小者
func MinIntAB(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//遍历数组且输出数值方法
func PrintArray(a []int) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d\t", a[i])
	}
	fmt.Println()
}

// 交换数组中的两个值
func Swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}
