package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"study/dataStructuresAndAlgorithms/dataStructures"
	"testing"
)

func TestReSort(t *testing.T) {
	m := 5
	n := 2
	A := []int{1, 3, 5, 7, 9, 2, 4, 6}
	fmt.Println("---Before resort---")
	for k, v := range A {
		fmt.Printf("index=%d\t value=%d\n", k, v)
	}
	// 重排序
	if ok := algorithms.ReSort(A, m, n); ok {
		fmt.Println("---After resort---")
		for k, v := range A {
			fmt.Printf("index=%d\t value=%d\n", k, v)
		}
	}
}

func TestSubLinkList(t *testing.T) {
	a := []int{1, 3, 5, 7, 9, 13, 17, 222}
	b := []int{3, 5, 8, 9, 77}

	var A = dataStructures.LNode{Next: nil, Data: len(a)}
	var B = dataStructures.LNode{Next: nil, Data: len(b)}
	A.CreateListR(a)
	B.CreateListR(b)

	A.PrintLinkList()
	B.PrintLinkList()

	// 链表求差集
	fmt.Println("---After sub---")
	if ok := algorithms.SubLinkList(&A, &B); ok {
		A.PrintLinkList()
	}
}

func TestArrayReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
	algorithms.ArrayReverse(a)
	fmt.Println("---After reverse---")
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func TestLinkListReverse(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := dataStructures.LNode{Next: nil, Data: len(data)}
	ln.CreateListR(data)
	fmt.Println("---Before LinkList Reverse---")
	ln.PrintLinkList()
	algorithms.LinkListReverse(&ln)
	fmt.Println("---After LinkList Reverse")
	ln.PrintLinkList()
}

func TestOnceQuickSort(t *testing.T) {
	a := []int{5, 8, 6, 4, 1, 3, 7, 2, 9, 10, 5, -1, 9, -6, 55, -3}
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
	algorithms.OnceQuickSort(a)
	fmt.Println("---After once quick sort---")
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

func TestGetMinByI(t *testing.T) {
	var data = []int{8, 7, 3, 4, 5, 6, 9}
	min := algorithms.GetMinByI(data)
	fmt.Println("---The original array is---")
	for _, v := range data {
		fmt.Printf("%d\t", v)
	}
	fmt.Printf("\nThe minimum is %d\n", min)
}

func TestIsMajority(t *testing.T) {
	a := []int{0, 5, 5, 3, 5, 7, 5, 5}
	b := []int{0, 5, 5, 3, 5, 1, 5, 7}

	ok1 := algorithms.IsMajority(a)
	fmt.Println("---The original array is---")
	for _, v := range a {
		fmt.Printf("%d\t", v)
	}
	if ok1 != -1 {
		fmt.Printf("\nThe majority of array exists and the value is 【%d】\n", ok1)
	} else {
		fmt.Printf("\nThe majority of array doesnot exists!\n")
	}

	ok2 := algorithms.IsMajority(b)
	fmt.Println("---The original array is---")
	for _, v := range b {
		fmt.Printf("%d\t", v)
	}
	if ok2 != -1 {
		fmt.Printf("\nThe majority of array exists and the value is %d\n", ok2)
	} else {
		fmt.Printf("\nThe majority of array doesnot exists!\n")
	}
}
