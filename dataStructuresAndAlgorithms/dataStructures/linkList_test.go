package dataStructures

import (
	"fmt"
	"testing"
)

func TestInitLinkListR(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := LNode{Next: nil, Data: len(data)}
	ln.CreateListR(data)
	p := &ln
	for {
		if p == nil {
			break
		}
		fmt.Println("LNode.data=", p.Data)
		p = p.Next
	}
}

func TestInitLinkListF(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := LNode{Next: nil, Data: len(data)}
	CreateListF(&ln, data)
	p := &ln
	for {
		if p == nil {
			break
		}
		fmt.Println("LNode.data=", p.Data)
		p = p.Next
	}
}

func TestFindAndDelete(t *testing.T) {
	var data = []int{1, 2, 3, 4, 5, 6}
	ln := LNode{Next: nil, Data: len(data)}
	ln.CreateListR(data)
	fmt.Println("---Before Delete---")
	p := &ln
	for {
		if p == nil {
			break
		}
		fmt.Println("LNode.data=", p.Data)
		p = p.Next
	}
	// 删除元素
	x := 8
	if ok := FindAndDelete(&ln, x); ok {
		fmt.Println("---After delete---")
		q := &ln
		for {
			if q == nil {
				break
			}
			fmt.Println("LNode.data=", q.Data)
			q = q.Next
		}
	} else {
		fmt.Printf("Cannot find the element=%d, delete failed!\n", x)
	}

}

func TestMerge(t *testing.T) {
	var data1 = []int{1, 3, 5, 7, 9, 10, 11, 15, 22}
	var data2 = []int{2, 4, 6, 8, 10, 13, 14, 16, 77}
	a := LNode{Next: nil, Data: len(data1)}
	a.CreateListR(data1)
	b := LNode{Next: nil, Data: len(data2)}
	b.CreateListR(data2)
	c := MergeOrderLinkList(&a, &b)
	p := c
	for {
		if p == nil {
			break
		}
		fmt.Println("LNode.data=", p.Data)
		p = p.Next
	}
}
