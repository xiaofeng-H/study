package dataStructures

import "fmt"

// 顺序队定义
type SequentialQueue struct {
	Data  [MaxSize]interface{}
	Front int
	Rear  int
}

/*顺序队（循环队列）操作 start*/

// 初始化队列
func (qu *SequentialQueue) initSqQu() {
	qu.Front = 0
	qu.Rear = 0
}

// 判断队空
func (qu *SequentialQueue) isEmptySqQu() bool {
	return qu.Front == qu.Rear
}

// 进队
func (qu *SequentialQueue) enSqQueue(data interface{}) bool {
	if (qu.Rear+1)%MaxSize == qu.Front {
		fmt.Println("The queue is full and cannot push any element into it")
		return false
	}

	qu.Rear = (qu.Rear + 1) % MaxSize
	qu.Data[qu.Rear] = data
	return true
}

// 出队
func (qu *SequentialQueue) deSqQueue() (bool, interface{}) {
	if qu.isEmptySqQu() {
		fmt.Println("The queue is empty and cannot get any element!")
		return false, nil
	}

	qu.Front = (qu.Front + 1) % MaxSize
	return true, qu.Data[qu.Front]
}

/*顺序队（循环队列）操作 end*/

/*链队定义 start*/

// 队结点类型定义
type QNode struct {
	Data interface{}
	Next *QNode
}

// 链队类型定义
type LinkQueue struct {
	Front *QNode
	Rear  *QNode
}

/*链队定义 end*/

/*链式队列操作 start*/

// 初始化队列
func (lq *LinkQueue) initLinkQu() {
	lq.Front = nil
	lq.Rear = nil
}

// 判断队空
func (lq *LinkQueue) isEmptyLinkQu() bool {
	return lq.Front == nil || lq.Rear == nil
}

// 进队
func (lq *LinkQueue) enLinkQueue(data interface{}) {
	// 构造新结点
	node := new(QNode)
	node.Data = data
	node.Next = nil
	// 尾插法
	// 注意：如果链队为空，则新插入的结点亦是队首结点
	if lq.isEmptyLinkQu() {
		lq.Front = node
		lq.Rear = node
	} else {
		lq.Rear.Next = node
		lq.Rear = node
	}
}

// 出队
func (lq *LinkQueue) deLinkQueue() (bool, interface{}) {
	if lq.isEmptyLinkQu() {
		fmt.Println("The queue is empty and cannot get any element!")
		return false, nil
	}

	// 出队
	data := lq.Front.Data
	// 如果出队之后队列为空，则同时要修改rear指针
	if lq.Front == lq.Rear {
		lq.initLinkQu()
	} else {
		lq.Front = lq.Front.Next
	}
	return true, data
}

/*链式队列操作 end*/
