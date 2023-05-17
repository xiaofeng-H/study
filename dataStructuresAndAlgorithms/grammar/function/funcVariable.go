package main

import (
	"fmt"
	"log"
)

// 使用头等函数实现斐波拉契数列
func fibonacciSequence(n int) int {
	if n <= 2 {
		log.Fatal("The required num is bigger than 2!!!")
	}
	t := tool()
	var res int
	for i := 2; i <= n; i++ {
		res = t()
		fmt.Printf("i=%d fibo=%d ", i, res)
	}
	fmt.Println()
	return res
}

// 求斐波拉契数列的前n个值
func tool() func() int {
	x0 := 0
	x1 := 1
	var x2 int
	return func() int {
		x2 = x0 + x1
		x0 = x1
		x1 = x2
		return x2
	}
}

// 函数值类型的使用，增加了代码的灵活性，类似于重载
// 案例：在某个业务场景下，某个接口有两个可能存在的功能，求两数之和或之差。
// 实现如下
// 定义可比较的泛型
type MyComparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | float32 | float64
}

type CalculateFunc func(a, b int) int

func Calculate(a, b int, calculate CalculateFunc) int {
	return calculate(a, b)
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
