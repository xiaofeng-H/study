package test

import (
	"fmt"
	"testing"
)

// 全排列
// 全排列结果集
var resPermute [][]int

func TestPermuteDemo(t *testing.T) {
	// 全排列数字
	var nums = []int{1, 2, 3, 4}
	// 记录已排列的数字
	var track = make([]int, 0)
	// 回溯处理
	backTrackPermute(nums, track)
	// 打印全排列结果
	for i := 0; i < len(resPermute); i++ {
		for j := 0; j < len(nums); j++ {
			fmt.Printf("%d \t", resPermute[i][j])
		}
		fmt.Println()
	}
	fmt.Printf("总共 %d 种排列\n", len(resPermute))
}

func backTrackPermute(nums, track []int) {
	// 边界值：如果已是全排列则记录
	if len(nums) == len(track) {
		tmp := make([]int, len(track))
		for i := 0; i < len(track); i++ {
			tmp[i] = track[i]
		}
		resPermute = append(resPermute, tmp)
	}

	// 回溯寻找全排列
	for i := 0; i < len(nums); i++ {
		// 检测当前元素是否可以加入排列
		if !canJoin(track, nums[i]) {
			continue
		}
		// 加入
		track = append(track, nums[i])
		// 回溯
		backTrackPermute(nums, track)
		// 剔除
		track = track[:len(track)-1]
	}
}

// 检测当前元素是否可以加入排列中
func canJoin(track []int, x int) bool {
	for _, v := range track {
		if v == x {
			return false
		}
	}
	return true
}

// N皇后
// N皇后结果集
var resNQueens [][][]byte

func TestNQueensDemo(t *testing.T) {
	// 规模
	var n = 8
	// 结果集
	resNQueens = make([][][]byte, 0)
	// 初始化棋盘
	var boards = make([][]byte, n)
	for j := 0; j < n; j++ {
		boards[j] = make([]byte, n)
		for i := 0; i < n; i++ {
			boards[j][i] = '.'
		}
	}
	// 回溯处理
	backTrackNQueens(boards, 0, n)
	// 打印结果
	for i := 0; i < len(resNQueens); i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				fmt.Printf("%s\t", string(resNQueens[i][j][k]))
			}
			fmt.Println()
		}
		fmt.Printf("=================第 %d 种解法=================\n", i+1)
	}
	fmt.Printf("%d 皇后问题共有 %d 种解\n", n, len(resNQueens))
}

func backTrackNQueens(boards [][]byte, row, n int) {
	// 边界值处理：如果已经处理完成则记录结果
	if row == n {
		tmp := make([][]byte, n)
		for i := 0; i < n; i++ {
			tmp[i] = make([]byte, n)
			for j := 0; j < n; j++ {
				tmp[i][j] = boards[i][j]
			}
		}
		resNQueens = append(resNQueens, tmp)
		return
	}

	// 寻找可以放皇后的位置并开始回溯
	for col := 0; col < n; col++ {
		if !isValid(boards, row, col) {
			continue
		}
		// 放置皇后
		boards[row][col] = 'Q'
		backTrackNQueens(boards, row+1, n)
		// 回溯
		boards[row][col] = '.'
	}
}

// 判断当前位置是否可以放皇后
func isValid(boards [][]byte, row, col int) bool {
	// 检查同一列是否已放置皇后
	for i := 0; i < row; i++ {
		if boards[i][col] == 'Q' {
			return false
		}
	}
	// 检查左上对角线是否已经放置皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if boards[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// 检查右上对角线是否已经放置皇后
	for i, j := row-1, col+1; i >= 0 && j < len(boards); {
		if boards[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
