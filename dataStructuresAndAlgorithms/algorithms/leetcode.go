package algorithms

import "strings"

/*
「力扣」第 53 题（最大子序和）
*/
// 动态规划
func MaxSubArray(nums []int) int {
	// 数组长度
	len := len(nums)
	// 定义状态： dp[i] 表示以i结尾的连续子序列的最大和
	var dp = make([]int, len, len)
	// 所求最大和
	var res int
	// 初始化状态
	dp[0] = nums[0]

	for i := 1; i < len; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	res = dp[0]
	for _, v := range dp {
		if res < v {
			res = v
		}
	}

	return res
}

/*
「力扣」第 70 题（爬楼梯）
*/
func climbStairs(n int) int {
	var f1, f2 = 1, 2
	var res int
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	for i := 3; i <= n; i++ {
		res = f1 + f2
		f1 = f2
		f2 = res
	}
	return res
}

/*
「力扣」第 118 题（杨辉三角1）
*/
func Generate(numRows int) [][]int {
	var res = make([][]int, numRows, numRows)
	res[0] = []int{1}
	if numRows >= 2 {
		res[1] = []int{1, 1}
	}

	for i := 2; i < numRows; i++ {
		res[i] = make([]int, i+1, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}

/*
「力扣」第 119 题（杨辉三角2）
*/
func GetRow(rowIndex int) []int {
	var pre, cur []int
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1)
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j]
		}
		pre = cur
	}
	return pre
}

/*
「力扣」第 2148 题（元素计数）
*/
func countElements(nums []int) int {
	// 满足同时具有严格较小值和严格较大值元素的数目
	var num int
	// 元素最小值
	min := int(1e5)
	// 最小值个数
	var minNum int
	// 元素最大值
	max := int(-1e5)
	// 最大值个数
	var maxNum int

	// 开始统计
	for i := 0; i < len(nums); i++ {
		// 统计最小值个数
		if nums[i] == min {
			minNum++
		} else if nums[i] < min {
			min = nums[i]
			minNum = 1
		}

		// 统计最大值个数
		if nums[i] == max {
			maxNum++
		} else if nums[i] > max {
			max = nums[i]
			maxNum = 1
		}
	}

	if len(nums) <= 2 {
		num = 0
	} else {
		if min == max {
			num = 0
		} else {
			num = len(nums) - minNum - maxNum
		}
	}

	return num
}

/*
「力扣」第 LCP17 题（速算机器人）
"A" 运算：使 x = 2 * x + y
"B" 运算：使 y = 2 * y + x
*/
func calculate(s string) int {
	x := 1
	y := 0
	for _, e := range s {
		if e == 'A' {
			x = operationA(x, y)
		} else if e == 'B' {
			y = operationB(x, y)
		}
	}

	return x + y
}

func operationA(x, y int) int {
	return 2*x + y
}

func operationB(x, y int) int {
	return 2*y + x
}

/*
「力扣」第 796 题（旋转字符串）
*/
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	ss := s + s
	return strings.Contains(ss, goal)
}

/*
「力扣」第 1108 题（IP地址无效化）
*/
func defangIPaddr(address string) string {
	split := strings.Split(address, ".")
	var res string
	for k, e := range split {
		if k == len(split)-1 {
			res = res + e
		} else {
			res = res + e + "[.]"
		}
	}
	return res
}

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/
/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*/
