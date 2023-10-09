package test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	var a = 1
	var p *int
	p = &a
	fmt.Printf("变量a的地址为：%p\n", &a)
	fmt.Printf("指针变量p为：%p\n", p)
	fmt.Printf("指针变量p的地址为：%p\n", &p)

	changePointer(&p)
	fmt.Printf("修改之后指针变量p为：%p\n", p)
	fmt.Printf("修改之后指针变量p的地址为：%p\n", &p)
}

func changePointer(p **int) {
	var b = 2
	fmt.Printf("修改函数中【修改前】指针变量p为：%p\n", *p)
	*p = &b
	fmt.Printf("修改函数中【修改后】指针变量p为：%p\n", *p)
}

func TestMake(t *testing.T) {
	var arr = make([][]int, 8)
	fmt.Printf("length:%d capacity:%d", len(arr), cap(arr))
	fmt.Printf("length:%d capacity:%d", len(arr[0]), cap(arr[0]))
}
