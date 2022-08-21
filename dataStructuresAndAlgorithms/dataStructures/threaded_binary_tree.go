package dataStructures

import (
	"fmt"
)

// 线索二叉树结点
type TBTNode struct {
	data interface{}
	// 标识域：0--指向孩子结点 | 1--线索
	lTag   byte
	rTag   byte
	lChild *TBTNode
	rChild *TBTNode
}

/* 中序线索二叉树---start--- */
// 二叉树初始化
func InitThreadBinaryTree(i int, arr []int) *TBTNode {
	if arr[i] == 0 {
		return nil
	}
	t := &TBTNode{arr[i], 0, 0, nil, nil}
	if i < len(arr) && 2*i+1 < len(arr) {
		t.lChild = InitThreadBinaryTree(2*i+1, arr)
	}
	if i < len(arr) && 2*i+2 < len(arr) {
		t.rChild = InitThreadBinaryTree(2*i+2, arr)
	}
	return t
}

var PERNOD *TBTNode = nil

// 二叉树中序线索化（递归算法）
func InThreadRec(p *TBTNode) {
	/*
		注意：线索化的时候每次指针赋值都要判空，不然会有空指针异常
	*/

	if p == nil {
		return
	}

	// 递归，左子树线索化（注意：递归左子树的前驱结点是pre而不是当前结点的父结点）
	InThreadRec(p.lChild)

	// 建立当前结点的前驱线索
	if p.lChild == nil {
		// 前驱线索化
		p.lTag = 1
		p.lChild = PERNOD
	}
	// 建立前驱结点的后继线索
	if PERNOD != nil && PERNOD.rChild == nil {
		PERNOD.rTag = 1
		PERNOD.rChild = p
	}
	// pre指向当前的node，作为node将要指向的下一个结点的前驱结点指示指针
	PERNOD = p
	// node指向下一个新结点，此时pre和node分别指向的结点形成了一个前驱后继对，为下一次线索的连接做准备
	p = p.rChild

	// 递归，右子树线索化
	InThreadRec(p)
}

// 建立中序线索二叉树
func CreateInThread(root *TBTNode) {
	if root == nil {
		fmt.Println("The BinaryTree is empty and cannot thread!")
		return
	}

	InThreadRec(root)
	PERNOD.rChild = nil
	PERNOD.rTag = 1
}

// 求以p为根的中序线索二叉树中，中序序列下的第一个结点
func GetFirstNode(p *TBTNode) *TBTNode {
	for p.lTag == 0 {
		p = p.lChild
	}
	return p
}

// 求在中序线索二叉树中，结点p在中序下的后继结点
func GetNextNode(p *TBTNode) *TBTNode {
	if p.rTag == 0 {
		return GetFirstNode(p.rChild)
	} else {
		// rTag == 1,直接返回后继线索
		return p.rChild
	}
}

// 以中序线索二叉树为存储结构的中序遍历
func InOrderByInOrderThreadTree(root *TBTNode) {
	// 建立中序二叉线索树
	CreateInThread(root)
	// 进行中序遍历
	for p := GetFirstNode(root); p != nil; p = GetNextNode(p) {
		visitBT(p.data)
	}
}

// 中序线索二叉树中序非递归遍历
func InOrderThreadedBTNoRec(root *TBTNode) {
	if root == nil {
		return
	}

	InOrderThreadedBTNoRec(root.lChild)

	visitBT(root.data)

	InOrderThreadedBTNoRec(root.rChild)
}

/* 中序线索二叉树---end--- */
