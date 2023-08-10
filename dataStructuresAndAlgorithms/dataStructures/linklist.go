package dataStructures

import (
	"fmt"
)

// @desc  	单链表结点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印单链表
func (ln *ListNode) PrintLinkList() {
	fmt.Printf("---The linkList`length is %d and data is:\n", ln.Val)
	p := ln.Next
	for {
		if p == nil {
			break
		}
		fmt.Printf("%d->", p.Val)
		p = p.Next
	}
	fmt.Print("nil\n")
}

// @desc  	        尾插法建立单链表
// @param  ln       单链表头结点
// @param  data     待插入的结点数据
func (ln *ListNode) CreateLinkListR(data []int) {
	if len(data) <= 0 {
		fmt.Println("The length of data is error!")
	}

	var s *ListNode // 用来指向新申请的结点
	var r *ListNode // 始终指向ln的终端结点

	r = new(ListNode)
	r = ln
	for i := 0; i < len(data); i++ {
		s = new(ListNode)
		s.Val = data[i]
		r.Next = s
		r = r.Next
	}
	r.Next = nil
}

// @desc  	        头插法建立单链表
// @param  ln       单链表头结点
// @param  data     待插入的结点数据
func CreateLinkListF(ln *ListNode, data []int) {
	if len(data) < 0 {
		fmt.Println("The length of data is error")
	}

	ln.Next = nil
	for i := 0; i < len(data); i++ {
		s := new(ListNode)
		s.Val = data[i]
		s.Next = ln.Next
		ln.Next = s
	}
}

// @desc  	        删除单链表中的数据为x的元素，成功返回true，失败返回false
// @param   ln      单链表的头结点
// @param   x       待删除的元素
// @return  bool    成功与否
func FindAndDelete(ln *ListNode, x int) bool {
	p := ln
	// 查找是否存在元素x
	for {
		if p.Next == nil {
			return false
		}
		if p.Next.Val == x {
			break
		}
		p = p.Next
	}
	// 找到则删除
	p.Next = p.Next.Next
	// 单链表长度-1
	ln.Val--
	return true
}

// @desc  	    将两个有序单链表（a,b）归并为一个有序单链表（c）
// @param   a   待merge的单链表
// @param   b   待merge的单链表
// @return  c   merge后得到的单链表
func MergeOrderLinkList(a, b *ListNode) (c *ListNode) {
	if a.Next == nil && b.Next == nil {
		fmt.Println("The list is both nil!")
	}
	/*
		指针指来指去的，特别容易出错，注意注意注意，Attention
		注意：归并为一个递增有序使用尾插法建立单链表；归并为一个递减有序使用头插法建立单链表。而不是改变判断条件大于为大于。
	*/
	p := a.Next
	q := b.Next
	c = a
	c.Val = a.Val + b.Val
	c.Next = nil // 必去要有指针先去接收a.next的结点
	fmt.Printf("The head is %v\n", c)
	r := new(ListNode)
	r = c
	for {
		if p == nil || q == nil {
			break
		}
		if p.Val <= q.Val {
			r.Next = p
			p = p.Next
			r = r.Next
		} else {
			r.Next = q
			q = q.Next
			r = r.Next
		}
	}

	if p != nil {
		r.Next = p
	}
	if q != nil {
		r.Next = q
	}

	return
}
