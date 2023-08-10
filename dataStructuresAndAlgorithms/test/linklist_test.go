package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"study/dataStructuresAndAlgorithms/dataStructures"
	"testing"
)

func TestInitLinkListR(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := dataStructures.ListNode{Next: nil, Val: len(data)}
	ln.CreateLinkListR(data)
	p := &ln
	for {
		if p == nil {
			break
		}
		fmt.Println("ListNode.Val=", p.Val)
		p = p.Next
	}
}

func TestInitLinkListF(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := dataStructures.ListNode{Next: nil, Val: len(data)}
	dataStructures.CreateLinkListF(&ln, data)
	p := &ln
	for {
		if p == nil {
			break
		}
		fmt.Println("ListNode.Val=", p.Val)
		p = p.Next
	}
}

func TestFindAndDelete(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := dataStructures.ListNode{Next: nil, Val: len(data)}
	ln.CreateLinkListR(data)
	fmt.Println("---Before Delete---")
	p := &ln
	for {
		if p == nil {
			break
		}
		fmt.Println("ListNode.Val=", p.Val)
		p = p.Next
	}
	// 删除元素
	x := 8
	if ok := dataStructures.FindAndDelete(&ln, x); ok {
		fmt.Println("---After delete---")
		q := &ln
		for {
			if q == nil {
				break
			}
			fmt.Println("ListNode.Val=", q.Val)
			q = q.Next
		}
	} else {
		fmt.Printf("Cannot find the element=%d, delete failed!\n", x)
	}

}

func TestMerge(t *testing.T) {
	var data1 = []int{1, 3, 5, 7, 9, 10, 11, 15, 22}
	var data2 = []int{2, 4, 6, 8, 10, 13, 14, 16, 77}
	a := dataStructures.ListNode{Next: nil, Val: len(data1)}
	a.CreateLinkListR(data1)
	b := dataStructures.ListNode{Next: nil, Val: len(data2)}
	b.CreateLinkListR(data2)
	c := dataStructures.MergeOrderLinkList(&a, &b)
	p := c
	for {
		if p == nil {
			break
		}
		fmt.Println("ListNode.Val=", p.Val)
		p = p.Next
	}
}

// 删除链表中值为X的所有元素
func TestDelX(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6, 5, 6, 7, 8, 5, 5, 5, 5, 5, 5, 8, 9, 88, 74, 4, 4, 5, 6}
	ln := dataStructures.ListNode{Next: nil, Val: len(data)}
	ln.CreateLinkListR(data)
	fmt.Println("原先的链表为：")
	ln.PrintLinkList()
	var x int = 5
	algorithms.DeleteXRec(&ln, x)
	fmt.Printf("删除%d之后的链表为：\n", x)
	ln.PrintLinkList()
}
