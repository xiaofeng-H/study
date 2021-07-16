package Algorithms

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
)

/*
算法描述：共享栈的设计，包括结构体和压栈出栈操作
算法要求：
算法规模：时间复杂度：O(1); 空间复杂度：O(1)
*/

// 共享栈结构体
type ShareStack struct {
	Data       [MaxSize]interface{}
	FrontIndex int // 共享栈头的栈顶
	RearIndex  int // 共享栈底的栈顶
}

// 压栈
func (ss *ShareStack) PushSS(data interface{}, isFront bool) bool {
	// 当前栈顶差值
	duration := ss.RearIndex - ss.FrontIndex
	if duration <= 0 {
		fmt.Println("The share stack is full and cannot push any element into it!")
		return false
	}

	// 共享栈头入栈
	if isFront {
		ss.FrontIndex++
		ss.Data[ss.FrontIndex] = data
	} else {
		ss.RearIndex--
		ss.Data[ss.RearIndex] = data
	}
	return true
}

// 出栈
func (ss *ShareStack) PullSS(isFront bool) (bool, interface{}) {
	// 判断不同栈底是否为空
	if isFront {
		if ss.FrontIndex == -1 {
			fmt.Println("The share stack is empty and cannot get any element!")
			return false, nil
		}
	} else {
		if ss.RearIndex == MaxSize {
			fmt.Println("The share stack is empty and cannot get any element!")
			return false, nil
		}
	}

	// 共享栈头出栈
	if isFront {
		data := ss.Data[ss.FrontIndex]
		ss.FrontIndex--
		return true, data
	} else {
		data := ss.Data[ss.RearIndex]
		ss.RearIndex++
		return true, data
	}
}

/*
算法描述：利用两个栈来实现队列的判断队列为空；入队及出队操作
算法要求：
算法规模：时间复杂度：O(1); 空间复杂度：O(1)
*/

// 队列结构体定义（使用两个顺序栈）
type QueueByDoubleStack struct {
	insStack dataStructures.SequentialStack
	delStack dataStructures.SequentialStack
}

// 判断队列是否为空
func (qs *QueueByDoubleStack) IsEmptyQS() bool {
	return qs.insStack.Top == -1 && qs.delStack.Top == -1
}

// 进队
func (qs *QueueByDoubleStack) InsertQS(data interface{}) bool {
	// 如果入栈已满但出栈不为空，则不能进队
	if qs.insStack.Top == MaxSize-1 && !qs.delStack.IsEmptySq() {
		fmt.Println("The insert stack is full but the delete stack is not empty! Cannot entry queue!")
		return false
	}

	// 入栈未满则进队
	if qs.insStack.Top < MaxSize {
		/*qs.insStack.Top++
		qs.insStack.Data[qs.insStack.Top] = data*/

		// 直接调用封装好的方法
		qs.insStack.PushSq(data)
	}
	// 入栈已满但出栈为空则进队
	if qs.insStack.Top == MaxSize-1 && qs.delStack.IsEmptySq() {
		/*// 转移元素
		for i := 0; i <= qs.insStack.Top; i++ {
			qs.delStack.Data[i] = qs.insStack.Data[i]
		}
		// 重置栈顶指针
		qs.insStack.Top = 0
		qs.delStack.Top = MaxSize - 1

		// 入队
		qs.insStack.Data[qs.insStack.Top] = data*/

		// 直接调用已封装好的方法
		// 转移元素
		for !qs.insStack.IsEmptySq() {
			if ok, data := qs.insStack.PopSq(); ok {
				qs.delStack.PushSq(data)
			}
		}
		// 入队
		qs.insStack.PushSq(data)
	}
	return true
}

// 出队
func (qs *QueueByDoubleStack) DeleteQS() (bool, interface{}) {
	// 入栈和出栈都为空时不可出栈
	if qs.IsEmptyQS() {
		fmt.Println("The insert stack is empty and the delete stack is empty!")
		return false, nil
	}

	// 如果出栈不为空，直接出队
	if !qs.delStack.IsEmptySq() {
		/*data := qs.delStack.Data[qs.delStack.Top]
		qs.delStack.Top--
		return true, data*/

		// 调用已经封装好的方法
		if ok, data := qs.delStack.PopSq(); ok {
			return true, data
		}
		return false, nil
	}
	// 如果出栈为空但入栈不为空则出队
	if qs.delStack.IsEmptySq() && !qs.insStack.IsEmptySq() {
		/*// 转移元素
		for i := 0; i < qs.insStack.Top; i++ {
			qs.delStack.Data[i] = qs.insStack.Data[i]
		}
		// 重置栈底指针
		qs.insStack.Top = -1
		qs.delStack.Top = lMaxSize - 2

		data := qs.delStack.Data[qs.delStack.Top]
		return true, data*/

		// 直接调用已经封装好的方法
		// 转移元素
		for !qs.insStack.IsEmptySq() {
			if ok, data := qs.insStack.PopSq(); ok {
				qs.delStack.PushSq(data)
			}
		}
		// 直接出队
		if ok, data := qs.delStack.PopSq(); ok {
			return true, data
		}
		return false, nil
	}
	return false, nil
}

/*
算法描述：循环队列的实现（设标记位tag：队列是否为空）
算法要求：
算法规模：时间复杂度：O(1); 空间复杂度：O(1)
*/

// 队列结构体定义
type QueueByTag struct {
	data  [MaxSize]interface{}
	front int
	rear  int
	tag   byte // 0：空 | 1：非空
}

// 队列初始化
func (qt *QueueByTag) InitQT() {
	qt.front = 0
	qt.rear = 0
	qt.tag = 0
}

// 判空
func (qt *QueueByTag) IsEmptyQT() bool {
	return qt.front == qt.rear && qt.tag == 0
}

// 判满
func (qt *QueueByTag) IsFullQT() bool {
	return qt.front == qt.rear && qt.tag == 1
}

// 入队
func (qt *QueueByTag) enQT(data interface{}) bool {
	if qt.IsFullQT() {
		fmt.Println("The queue is full!")
		return false
	}

	qt.rear = (qt.rear + 1) % MaxSize
	qt.data[qt.rear] = data
	// 每次入队都重置tag值
	qt.tag = 1
	return true
}

// 出队
func (qt *QueueByTag) DelQT() (bool, interface{}) {
	if qt.IsEmptyQT() {
		fmt.Println("The queue is empty!")
		return false, nil
	}

	qt.front = (qt.front + 1) % MaxSize
	data := qt.data[qt.front]
	/*
		// 删除元素后如果队列空，则修改tag值
		if qt.rear == qt.front {
			qt.tag = 1
		}*/

	// 只要有元素出队，就重置tag
	qt.tag = 0
	return true, data
}

/*
算法描述：数值转换
算法要求：
算法规模：时间复杂度：O(logN + m); 空间复杂度：O(1)
*/
func BaseTrans(N int, B int) int {
	stack := [10000]int{}
	top := -1
	var res int
	for N != 0 {
		a := N % B
		N /= B
		top++
		stack[top] = a
	}
	for top != -1 {
		a := stack[top]
		top--
		res = res*10 + a
	}
	return res
}
