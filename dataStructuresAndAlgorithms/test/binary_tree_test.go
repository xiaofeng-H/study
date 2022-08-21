package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
	"testing"
)

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
