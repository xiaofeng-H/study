package dataStructures

import "math"

// 平衡二叉搜索树（Self-balancing binary search tree）又被称为AVL树（有别于AVL算法），且具有以下性质：
// 它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。

// E AVlTree 是 BST，所以结点必须是可比较的
type E int32

// Node AVL 结点...
type Node struct {
	e      E
	left   *Node
	right  *Node
	height int32
}

// NewNode 初始化一个结点
func NewNode(e E) *Node {
	return &Node{e: e,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

// GetHeight
// @desc  获取某一结点的高度
// @param  node
// @return  uint32
func GetHeight(node *Node) int32 {
	if node == nil {
		return 0
	}
	return node.height
}

// GetBalanceFactor
// @desc  获取结点的平衡因子
// @param  node
// @return  uint32
func GetBalanceFactor(node *Node) int32 {
	if node == nil {
		return 0
	}
	return GetHeight(node.right) - GetHeight(node.left)
}

// IsBalanced
// @desc  判断某个结点是否平衡
// @param  node
// @return  bool
func IsBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	if math.Abs(float64(GetBalanceFactor(node))) > 1 {
		return false
	}
	return IsBalanced(node.left) && IsBalanced(node.right)
}
