package algorithms

import (
	"fmt"
	"math"
)

/**
 * NC78 反转链表
 * @param pHead ListNode类
 * @return ListNode类
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	// 空链表处理
	if pHead == nil {
		return nil
	}
	// 使用头插法建立链表来实现链表的反转
	var head *ListNode
	head = nil
	for pHead != nil {
		p := pHead.Next
		pHead.Next = head
		head = pHead
		pHead = p
	}
	return head
}

// 明明的随机数
func Random() {
	// 接收数据1
	var N int
	fmt.Scan(&N)
	var arr []int
	for i := 0; i < N; i++ {
		var n int
		fmt.Scan(&n)
		arr = append(arr, n)
	}
	// 处理数据
	// 去重并排序（使用散列表思想）
	var hash [500]int = [500]int{0}
	for _, v := range arr {
		index := v % 500
		if hash[index] != 0 {
			continue
		}
		hash[index] = v
	}
	for _, v := range hash {
		if v != 0 {
			fmt.Println(v)
		}
	}
}

// 进制转换
func Conversion() {
	var strArr []string
	for {
		var str string
		n, err := fmt.Scanln(&str)
		if n != 0 {
			strArr = append(strArr, str)
		}
		if err != nil {
			break
		}
	}
	for _, v := range strArr {
		var res int = 0
		data := []byte(v)
		len := len(data)
		for i := 2; i < len; i++ {
			j := 0
			if data[i] == 'A' {
				j = 10
			} else if data[i] == 'B' {
				j = 11
			} else if data[i] == 'C' {
				j = 12
			} else if data[i] == 'D' {
				j = 13
			} else if data[i] == 'E' {
				j = 14
			} else if data[i] == 'F' {
				j = 15
			} else {
				j = int(data[i])
			}
			res = res + j*int(math.Pow(16, float64(len-i-1)))
		}
		fmt.Println(res)
	}
}

/**
 * NC39 N皇后问题
 * N 皇后问题是指在 n * n 的棋盘上要摆 n 个皇后，
 * 要求：任何两个皇后不同行，不同列也不在同一条斜线上，
 * 求给一个整数 n ，返回 n 皇后的摆法数。
 * 数据范围: 1 \le n \le 91≤n≤9
 * 要求：空间复杂度 O(1)O(1) ，时间复杂度 O(n!)O(n!)
 */
var nQueensRes [][]string

// 输入棋盘边长 n，返回所有合法的位置
func SolveNQueens(n int) [][]string {
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘
	var board [][]rune = make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	backTrack(board, 0)
	return nQueensRes
}

// 路径：board中小于row的那些行都已经成功放置了皇后
// 选择列表：第row行的所有列都是放置皇后的选择
// 结束条件：row超过board的最后一行
func backTrack(board [][]rune, row int) {
	// 触发结束条件
	if row == len(board) {
		var str []string
		for i := range board {
			tmp := string(board[i])
			str = append(str, tmp)
		}
		nQueensRes = append(nQueensRes, str)
		return
	}

	n := len(board[row])
	for col := 0; col < n; col++ {
		// 排除不合法选择
		if !isValid(board, row, col) {
			continue
		}
		// 做选择
		board[row][col] = 'Q'
		// 进入下一行决策
		backTrack(board, row+1)
		// 撤销选择
		board[row][col] = '.'
	}
}

// 是否可以在board[row][col]放置皇后
func isValid(board [][]rune, row int, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i < n; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	i := row - 1
	j := col + 1
	for i >= 0 && j < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	// 检查左上方是否有皇后互相冲突
	k := row - 1
	l := col - 1
	for k >= 0 && l >= 0 {
		if board[k][l] == 'Q' {
			return false
		}
		k--
		l--
	}
	return true
}


func threeOrders(root *TreeNode) [][]int {
	var res [][]int
	var resPre, resIn, resPost []int
	// write code here
	res = make([][]int, 0, 3)
	resPre = make([]int, 0)
	resIn = make([]int, 0)
	resPost = make([]int, 0)

	preOrder(root, &resPre)
	inOrder(root, &resIn)
	postOrder(root, &resPost)

	res = append(res, resPre)
	res = append(res, resIn)
	res = append(res, resPost)

	return res
}

func preOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preOrder(root.Left, res)
	preOrder(root.Right, res)
}

func inOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, res)
	*res = append(*res, root.Val)
	inOrder(root.Right, res)
}

func postOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, res)
	postOrder(root.Right, res)
	*res = append(*res, root.Val)
}
