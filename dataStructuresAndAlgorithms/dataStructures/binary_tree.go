package dataStructures

import (
	"fmt"
)

// 二叉树结点
type BTNode struct {
	data   interface{}
	lChild *BTNode
	rChild *BTNode
}

// 二叉树结点初始化
func NewBTNode(data interface{}) BTNode {
	return BTNode{
		data:   data,
		lChild: nil,
		rChild: nil,
	}
}

/* 二叉树初始化---start--- */
// 1.由标明空子树的先根遍历序列建立一颗二叉树的操作算法
var i = -1 // 结点数据段数组下标
func PreOrderInitRec(arr []int) *BTNode {
	i = i + 1
	if i >= len(arr) {
		return nil
	}
	var t BTNode
	if arr[i] != 0 {
		t = BTNode{arr[i], nil, nil}
		t.lChild = PreOrderInitRec(arr)
		t.rChild = PreOrderInitRec(arr)
	} else {
		return nil
	}
	return &t
}

// 2.由完全二叉树的顺序存储结构建立其二叉链式存储结构（下标从0开始，数字0表示该下标数组值为空，也即该结点为空）
func ArrayToLinkInitRec(i int, arr []int) *BTNode {
	if arr[i] == 0 {
		return nil
	}
	t := &BTNode{arr[i], nil, nil}
	if i < len(arr) && 2*i+1 < len(arr) {
		t.lChild = ArrayToLinkInitRec(2*i+1, arr)
	}
	if i < len(arr) && 2*i+2 < len(arr) {
		t.rChild = ArrayToLinkInitRec(2*i+2, arr)
	}
	return t
}

// 由先根遍历序列和中根遍历序列建立一颗二叉树
func PreMidInitRec(pa []int, ma []int, preI int, midI int, len int) *BTNode {
	if len > 0 {
		p := pa[preI]
		i := 0
		for ; i < len; i++ {
			if p == ma[midI+i] {
				break
			}
		}
		t := BTNode{p, nil, nil}
		t.lChild = PreMidInitRec(pa, ma, preI+1, midI, i)
		t.rChild = PreMidInitRec(pa, ma, preI+i+1, midI+i+1, len-1-i)
		return &t
	}
	return nil
}

/* 二叉树初始化---end--- */

/* 二叉树遍历---start--- */
// 二叉树先根遍历（递归）
func PreOrderRec(bt *BTNode) {
	if bt == nil {
		return
	}
	visitBT(bt.data)
	PreOrderRec(bt.lChild)
	PreOrderRec(bt.rChild)
}

// 二叉树先根遍历（非递归）
func PreOrderNoRec(bt *BTNode) {
	if bt == nil {
		fmt.Println("The BinaryTree is empty!")
		return
	}

	// 定义一个栈
	var stack []*BTNode = make([]*BTNode, MaxSize)
	top := -1
	var element *BTNode // 栈顶元素
	// 1.先让头结点进栈
	top++
	stack[top] = bt
	// 2.栈不空的时候进行遍历
	for top != -1 {
		// 3.栈顶元素出栈，结点右孩子先入栈，然后左孩子再入栈，依次循环
		element = stack[top]
		top--
		visitBT(element.data)
		if element.rChild != nil {
			top++
			stack[top] = element.rChild
		}
		if element.lChild != nil {
			top++
			stack[top] = element.lChild
		}
	}
}

// 二叉树中根遍历（递归）
func InOrderRec(bt *BTNode) {
	if bt == nil {
		return
	}
	InOrderRec(bt.lChild)
	visitBT(bt.data)
	InOrderRec(bt.rChild)
}

// 二叉树中根遍历（非递归）
func InOrderNoRec(bt *BTNode) {
	if bt == nil {
		fmt.Println("The BinaryTree is empty!")
		return
	}

	// 定义一个栈
	var stack []*BTNode = make([]*BTNode, MaxSize)
	top := -1
	var element *BTNode // 被操作元素
	// 1.先让被操作元素指向头结点
	element = bt
	// 2.栈不空或者被操作元素不为空的时候进行遍历
	for top != -1 || element != nil {
		// 3.让被操作元素及自己所有的左子孙入栈，直到没有左子孙
		for element != nil {
			top++
			stack[top] = element
			element = element.lChild
		}
		// 4.栈不空的情况下出栈，并让被操作元素指向出栈元素的右孩子，以此来让被操作元素右孩子的所有左子孙入栈，
		// 依次循环，便可以按中根遍历遍历完所有的元素
		if top != -1 {
			element = stack[top]
			top--
			visitBT(element.data)
			// 被操作元素指向自己的右孩子
			element = element.rChild
		}
	}
}

// 二叉树后根遍历（递归）
func PostOrderRec(bt *BTNode) {
	if bt == nil {
		return
	}
	PostOrderRec(bt.lChild)
	PostOrderRec(bt.rChild)
	visitBT(bt.data)
}

// 二叉树后根遍历（非递归）
/*
	注意：观察可得，先将后序遍历得到的后序序列逆序得到逆后序序列，不难发现，逆后序遍历序列只不过是先序遍历过程中
对左右子树遍历顺序交换得到的结果，所以在先序非递归遍历的算法上稍稍加工便可（多用一个栈来保存逆后序遍历序列）。
*/
func PostOrderNoRec(bt *BTNode) {
	if bt == nil {
		fmt.Println("The BinaryTree is empty!")
		return
	}

	// 定义两个栈
	var stack1 []*BTNode = make([]*BTNode, MaxSize) // 实现逆后序遍历过程（辅助栈）
	var stack2 []*BTNode = make([]*BTNode, MaxSize) // 保存逆后序遍历的结果（保存栈）
	top1 := -1
	top2 := -1
	var element *BTNode // 栈顶元素
	// 1.先让头结点进入辅助栈
	top1++
	stack1[top1] = bt
	// 2.辅助栈不空的时候进行遍历
	for top1 != -1 {
		// 3.辅助栈栈顶元素出栈，并将出栈元素压入保存栈
		element = stack1[top1]
		top1--
		top2++
		stack2[top2] = element
		// 4.让栈顶元素的左孩子先入辅助栈，再让右孩子入辅助栈，此过程使得得到的出栈序列为逆后序序列，
		// 是与先序遍历非递归算法最大的区别
		if element.lChild != nil {
			top1++
			stack1[top1] = element.lChild
		}
		if element.rChild != nil {
			top1++
			stack1[top1] = element.rChild
		}
	}
	// 5.以上过程得到了逆后序遍历序列，并且保存到了保存栈中，此处只需遍历保存栈便可得到后序遍历序列
	for top2 != -1 {
		element = stack2[top2]
		top2--
		visitBT(element.data)
	}
}

// 二叉树层次遍历（非递归）
func LevelTraverse(bt *BTNode) {
	// 循环队列头尾指针
	front := int32(0)
	rear := int32(0)
	// 定义一个循环队列，用来记录将要访问的层次上的结点
	var que [MaxSize]*BTNode

	if bt != nil {
		// 若队列已满，则层次遍历不能继续进行，需要更大容量的队列
		if (rear+1)%MaxSize == front {
			fmt.Print("The queue is full and cannot push any element into it!")
			return
		}

		rear = (rear + 1) % MaxSize
		que[rear] = bt // 根结点入队
		// 当队列不空的时候进行循环
		for front != rear {
			front = (front + 1) % MaxSize
			q := que[front]
			// 访问对头结点，此处只简单打印
			visitBT(q.data)
			// 如果左子树不空，则左子树的根结点入队
			if q.lChild != nil {
				// 此处假设队列容量大于结点个数，不做队列满判断，下同
				rear = (rear + 1) % MaxSize
				que[rear] = q.lChild
			}
			// 如果右子树不空，则左子树的根结点入队
			if q.rChild != nil {
				// 此处假设队列容量大于结点个数，不做队列满判断，下同
				rear = (rear + 1) % MaxSize
				que[rear] = q.rChild
			}
		}
	}
}

/* 二叉树遍历---end--- */

// 操作遍历元素（此处只做简单打印）
func visitBT(data interface{}) {
	fmt.Printf("%v, ", data)
}
