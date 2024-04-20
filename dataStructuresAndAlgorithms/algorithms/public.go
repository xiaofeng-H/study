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

// 遍历数组且输出数值方法
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
