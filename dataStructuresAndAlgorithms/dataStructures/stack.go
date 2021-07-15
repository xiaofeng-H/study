package dataStructures

import "fmt"

// 顺序栈定义
type SequentialStack struct {
	Data []interface{}
	Top  int
}

/*顺序栈操作 start*/

// 初始化
func (ss *SequentialStack) initSq() {
	ss.Top = -1
}

// 判断栈是否为空
func (ss *SequentialStack) isEmptySq() bool {
	return ss.Top == -1
}

// 压栈
func (ss *SequentialStack) pushSq(data interface{}) bool {
	// 栈满则不能继续入栈
	if ss.Top == len(ss.Data)-1 {
		fmt.Println("The stack is full and cannot push any element into it!")
		return false
	}

	// 入栈
	ss.Top++
	ss.Data[ss.Top] = data
	return true
}

// 出栈
func (ss *SequentialStack) popSq() (bool, interface{}) {
	// 栈空则不能出栈
	if ss.Top == -1 {
		fmt.Println("The stack is empty and cannot pop any element!")
		return false, nil
	}

	// 出栈
	data := ss.Data[ss.Top]
	ss.Top--
	return true, data

}

/*顺序栈操作 end*/

// 链栈结点定义
type LinkStack struct {
	Data interface{}
	Next *LinkStack
}

/*带头结点的链栈操作 start*/

// 初始化
func (ls *LinkStack) initLs() {
	ls.Next = nil
}

// 判断栈是否为空
func (ls *LinkStack) isEmptyLs() bool {
	return ls.Next == nil
}

// 压栈（不存在失败的情况，除非内存不足）
func (ls *LinkStack) pushLs(data interface{}) {
	// 头插法
	node := new(LinkStack)
	node.Data = data
	node.Next = ls.Next
	ls.Next = node
}

// 出栈
func (ls *LinkStack) popLs() (bool, interface{}) {
	if ls.Next == nil {
		fmt.Println("The stack is empty and cannot pop any element!")
		return false, nil
	}

	data := ls.Next.Data
	ls.Next = ls.Next.Next
	return true, data
}

/*带头结点的链栈操作 end*/

/*不带头结点的链栈操作 start*/

// 初始化
func (ls *LinkStack) initLsl() {
	ls = nil
}

// 判断栈是否为空
func (ls *LinkStack) isEmptyLsl() bool {
	return ls == nil
}

// 压栈（不存在失败的情况，除非内存不足）
func (ls *LinkStack) pushLsl(data interface{}) {
	// 头插法
	node := new(LinkStack)
	node.Data = data
	node.Next = ls
	ls = node
}

// 出栈
func (ls *LinkStack) popLsl() (bool, interface{}) {
	if ls == nil {
		fmt.Println("The stack is empty and cannot pop any element!")
		return false, nil
	}

	data := ls.Data
	ls = ls.Next
	return true, data
}

/*不带头结点的链栈操作 start*/
