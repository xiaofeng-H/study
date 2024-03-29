package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
	"testing"
)

// 二叉树遍历
func TestBinaryTree(t *testing.T) {
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var arr1 = []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	var arr2 = []int{1, 2, 3, 4, 5, 6, 7, 0, 0, 0, 0, 8, 9}
	//bt := PreOrderInitRec(arr)	// 先根遍历方式初始化二叉树
	bt := dataStructures.ArrayToLinkInitRec(0, arr2) // 顺序结构转为链式结构
	fmt.Println("---先根遍历二叉树结点（递归）---")
	dataStructures.PreOrderRec(bt)
	fmt.Println()
	fmt.Println("---先根遍历二叉树结点（非递归）---")
	dataStructures.PreOrderNoRec(bt)
	fmt.Println()
	fmt.Println("---中根遍历二叉树结点（递归）---")
	dataStructures.InOrderRec(bt)
	fmt.Println()
	fmt.Println("---中根遍历二叉树结点（非递归）---")
	dataStructures.InOrderNoRec(bt)
	fmt.Println()
	fmt.Println("---后根遍历二叉树结点（递归）---")
	dataStructures.PostOrderRec(bt)
	fmt.Println()
	fmt.Println("---后根遍历二叉树结点（非递归）---")
	dataStructures.PostOrderNoRec(bt)
	fmt.Println()
	fmt.Println("---层次遍历二叉树结点---")
	dataStructures.LevelTraverse(bt)
	fmt.Println()
}

// 中序线索二叉树
func TestInOrderByInOrderThreadTree(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bt := dataStructures.InitThreadBinaryTree(0, arr) // 顺序结构转为链式结构
	fmt.Println("---非递归中序遍历:")
	dataStructures.InOrderThreadedBTNoRec(bt)
	fmt.Println()
	fmt.Println("---中序线索二叉树遍历:")
	dataStructures.InOrderByInOrderThreadTree(bt)
	fmt.Println()
}

// 前序中序创建二叉树
func TestPreInInitBT(t *testing.T) {
	var pre = []rune{'A', 'B', 'C', 'D', 'E', 'F'}
	var in = []rune{'C', 'B', 'A', 'E', 'D', 'F'}
	bt := dataStructures.PreInInitBT(pre, in, 0, len(pre)-1, 0, len(in)-1)
	dataStructures.PostOrderNoRec(bt)
}

func TestInitSpecialBT(t *testing.T) {
	bt := dataStructures.InitBTByPreInOrder()
	//dataStructures.PostOrderNoRec(bt)
	preOrder(bt)
}

// 求二叉树最大深度（递归思想理解）
var res int = 0
var depth int = 0

func preOrder(bt *dataStructures.BTNode) {
	if bt == nil {
		return
	}
	depth++
	fmt.Println(depth)
	if bt.LChild == nil && bt.RChild == nil {
		if res < depth {
			res = depth
		}
	}
	preOrder(bt.LChild)
	preOrder(bt.RChild)
	depth--
	fmt.Println(depth)
}

// 求二叉树值为x的结点所在的层号
func TestGetLevelBTN(t *testing.T) {
	// m=5,n=6
	bt := dataStructures.InitBTByPreInOrder()
	x := 'N'
	dataStructures.GetLevelOfBTN(bt, x)
}
