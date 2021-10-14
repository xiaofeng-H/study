package dataStructures

import (
	"fmt"
)

// @desc  	单链表结点定义
type LNode struct {
	Data int
	Next *LNode
}

// 打印单链表
func (ln *LNode) PrintLinkList() {
	fmt.Printf("---The linkList`length is %d and data is:\n", ln.Data)
	p := ln.Next
	for {
		if p == nil {
			break
		}
		fmt.Printf("%d->", p.Data)
		p = p.Next
	}
	fmt.Print("nil\n")
}

// @desc  	        尾插法建立单链表
// @param  ln       单链表头结点
// @param  data     待插入的结点数据
func (ln *LNode) CreateListR(data []int) {
	if len(data) <= 0 {
		fmt.Println("The length of data is error!")
	}

	var s *LNode // 用来指向新申请的结点
	var r *LNode // 始终指向ln的终端结点

	r = new(LNode)
	r = ln
	for i := 0; i < len(data); i++ {
		s = new(LNode)
		s.Data = data[i]
		r.Next = s
		r = r.Next
	}
	r.Next = nil
}

// @desc  	        头插法建立单链表
// @param  ln       单链表头结点
// @param  data     待插入的结点数据
func CreateListF(ln *LNode, data []int) {
	if len(data) < 0 {
		fmt.Println("The length of data is error")
	}

	ln.Next = nil
	for i := 0; i < len(data); i++ {
		s := new(LNode)
		s.Data = data[i]
		s.Next = ln.Next
		ln.Next = s
	}
}

// @desc  	        删除单链表中的数据为x的元素，成功返回true，失败返回false
// @param   ln      单链表的头结点
// @param   x       待删除的元素
// @return  bool    成功与否
func FindAndDelete(ln *LNode, x int) bool {
	p := ln
	// 查找是否存在元素x
	for {
		if p.Next == nil {
			return false
		}
		if p.Next.Data == x {
			break
		}
		p = p.Next
	}
	// 找到则删除
	p.Next = p.Next.Next
	// 单链表长度-1
	ln.Data--
	return true
}

// @desc  	    将两个有序单链表（a,b）归并为一个有序单链表（c）
// @param   a   待merge的单链表
// @param   b   待merge的单链表
// @return  c   merge后得到的单链表
func MergeOrderLinkList(a, b *LNode) (c *LNode) {
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
	c.Data = a.Data + b.Data
	c.Next = nil // 必去要有指针先去接收a.next的结点
	fmt.Printf("The head is %v\n", c)
	r := new(LNode)
	r = c
	for {
		if p == nil || q == nil {
			break
		}
		if p.Data <= q.Data {
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
