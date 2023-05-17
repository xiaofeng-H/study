package main

import "fmt"

func main() {
	/*// 头等函数实现斐波拉契数列测试
	var n int = 5
	res := fibonacciSequence(n)
	fmt.Println(res)*/

	// 函数值类型案例测试
	var a, b int = 2, 3
	res1 := Calculate(a, b, sum)
	res2 := Calculate(a, b, sub)
	fmt.Println(res1, res2)
}
