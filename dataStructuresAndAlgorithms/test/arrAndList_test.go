package test

import (
	"fmt"
	"testing"
)

func TestMergeArr(t *testing.T) {
	var a = []int{1, 3, 5, 7, 9, 10, 11, 12}
	var b = []int{2, 4, 6, 8}
	res := merge(a, b)
	fmt.Println(res)
}
func merge(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	var sum []int = make([]int, 0, len(a)+len(b))
	for len(a) != 0 && len(b) != 0 {
		if a[0] < b[0] {
			sum = append(sum, a[0])
			a = a[1:]
		} else {
			sum = append(sum, b[0])
			b = b[1:]
		}
	}
	if len(a) == 0 {
		sum = append(sum, b...)
	} else {
		sum = append(sum, a...)
	}
	return sum
}

func TestReveList(t *testing.T) {
	var a = []int{1, 3, 5, 7, 9, 10, 11, 12}
	list := createList(a)
	p := list
	for p != nil {
		fmt.Printf("%v ", p.val)
		p = p.next
	}
	fmt.Println()
	res := reverseList(list)
	q := res
	for q != nil {
		fmt.Printf("%v ", q.val)
		q = q.next
	}
}

type listNode struct {
	val  int
	next *listNode
}

func createList(val []int) *listNode {
	var list = &listNode{}
	var p = list
	for _, v := range val {
		tmp := &listNode{
			val:  v,
			next: nil,
		}
		p.next = tmp
		p = p.next
	}
	return list.next
}

func reverseList(ls *listNode) *listNode {
	if ls == nil {
		fmt.Println("The list is nil!")
		return nil
	}
	p := &listNode{}
	for ls != nil {
		q := ls.next
		ls.next = p.next
		p.next = ls
		ls = q
	}
	return p.next
}

//实现一个函数，输入一个整数n，输出从1到n的所有整数中，3的倍数用"Fizz"表示，5的倍数用"Buzz"表示，既是3的倍数又是5的倍数的数用"FizzBuzz"表示，其余数直接输出。
//要求：
//1. 输入参数n为正整数，并且n>=1；
//2. 函数输出一个字符串切片，包含从1到n的所有整数的表示；
//3. 如果一个数是3的倍数，用"Fizz"表示；
//4. 如果一个数是5的倍数，用"Buzz"表示；
//5. 如果一个数既是3的倍数又是5的倍数，用"FizzBuzz"表示；
//6. 其余数直接输出。
//例如，当n=15时，输出如下：
//["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
func TestConvert(t *testing.T) {
	res := convert(15)
	fmt.Println(res)
}
func convert(n int) []string {
	if n == 0 {
		return nil
	}
	var res []string = make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			res = append(res, "Fizz")
			continue
		}
		if i%5 == 0 {
			res = append(res, "Buzz")
			continue
		}
		res = append(res, fmt.Sprint(i))
	}
	return res
}
