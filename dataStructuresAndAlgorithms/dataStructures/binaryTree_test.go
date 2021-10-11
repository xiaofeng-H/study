package dataStructures

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	//var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//var arr1 = []int{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	var arr2 = []int{1, 2, 3, 4, 5, 6, 7, 0, 0, 0, 0, 8, 9}
	//bt := PreOrderInitRec(arr)	// 先根遍历方式初始化二叉树
	bt := ArrayToLinkInitRec(0, arr2) // 顺序结构转为链式结构
	fmt.Println("---先根遍历二叉树结点（递归）---")
	PreOrderRec(bt)
	fmt.Println()
	fmt.Println("---先根遍历二叉树结点（非递归）---")
	PreOrderNoRec(bt)
	fmt.Println()
	fmt.Println("---中根遍历二叉树结点（递归）---")
	InOrderRec(bt)
	fmt.Println()
	fmt.Println("---中根遍历二叉树结点（非递归）---")
	InOrderNoRec(bt)
	fmt.Println()
	fmt.Println("---后根遍历二叉树结点（递归）---")
	PostOrderRec(bt)
	fmt.Println()
	fmt.Println("---后根遍历二叉树结点（非递归）---")
	PostOrderNoRec(bt)
	fmt.Println()
	fmt.Println("---层次遍历二叉树结点---")
	LevelTraverse(bt)
	fmt.Println()
}
