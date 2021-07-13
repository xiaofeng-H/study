package Algorithms

import (
	"fmt"
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
	if ok := ReSort(A, m, n); ok {
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
	if ok := SubLinkList(&A, &B); ok {
		A.PrintLinkList()
	}
}
