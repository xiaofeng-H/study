package algorithms

import (
	"math"
	"strings"
)

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
func ClimbStairs(n int) int {
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
func CountElements(nums []int) int {
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
func Calculate(s string) int {
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
func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	ss := s + s
	return strings.Contains(ss, goal)
}

/*
「力扣」第 1108 题（IP地址无效化）
*/
func DefangIPaddr(address string) string {
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

/*======================================滑动串口start============================================*/
/*
「力扣」第 76 题（最小覆盖子串）
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：
	对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
	如果 s 中存在这样的子串，我们保证它是唯一的答案。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
滑动窗⼝算法的思路是这样：
1、我们在字符串  S  中使⽤双指针中的左右指针技巧，初始化  left =
right = 0  ，把索引左闭右开区间  [left, right)  称为⼀个「窗⼝」。
2、我们先不断地增加  right  指针扩⼤窗⼝  [left, right)  ，直到窗⼝中
的字符串符合要求（包含了  T  中的所有字符）。
3、此时，我们停⽌增加  right  ，转⽽不断增加  left  指针缩⼩窗⼝
[left, right)  ，直到窗⼝中的字符串不再符合要求（不包含  T  中的所有
字符了）。同时，每次增加  left  ，我们都要更新⼀轮结果。
4、重复第 2 和第 3 步，直到  right  到达字符串  S  的尽头。
*/
func MinWindow(s string, t string) string {
	// 首先，初始化window和need两个哈希表，记录窗口中的字符和需要凑齐的字符
	var need, window map[byte]int
	need = make(map[byte]int)
	window = make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	// 然后，使用left和right变量初始化窗口的两端，不要忘了，区间[left,right)
	// 是左闭右开的，所以初始情况下窗口没有包含任何元素
	left := 0
	right := 0
	// valid变量表示窗口中满足need条件的字符个数，如果valid和len(need)的大小相同，
	// 则说明窗口已满足条件，已经完全覆盖了串t
	valid := 0
	// 记录最小覆盖子串的起始索引及长度
	start := 0
	length := math.MaxInt32
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		for valid == len(need) {
			// 在这里更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			// d 是将移出窗口的字符
			d := s[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	} else {
		return s[start : start+length]
	}
}

/*
「力扣」第 567 题（字符串的排列）
给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
换句话说，s1 的排列之一是 s2 的 子串 。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-in-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func CheckInclusion(s1 string, s2 string) bool {
	// 初始化数据
	need := make(map[byte]int)
	window := make(map[byte]int)
	for e := range s1 {
		need[s1[e]]++
	}
	var left, right, valid int = 0, 0, 0

	for right < len(s2) {
		c := s2[right]
		// 右移窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 判断左侧窗口是否要收缩
		for right-left >= len(s1) {
			// 在这里判断是否找到了合法的子串
			if valid == len(need) {
				return true
			}
			d := s2[left]
			// 左移窗口
			left++
			// 进行窗口内数据的一系列更新
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	// 未找到符合条件的子串
	return false
}

/*
「力扣」第 438 题（找到字符串中所有字母异位词）
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func FindAnagrams(s string, p string) []int {
	// 初始化数据
	window := make(map[byte]int)
	need := make(map[byte]int)
	for k := range p {
		need[p[k]]++
	}
	left, right, valid := 0, 0, 0
	// 所求子串的起始索引结果集
	res := make([]int, 0)

	// 滑动窗口
	for right < len(s) {
		c := s[right]
		// 右移窗口并进行一系列数据的更新
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否需要左移
		for right-left >= len(p) {
			// 当窗口符合条件时，把起始索引加入res
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			// 左移窗口并进行一系列的更新
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

/*
「力扣」第 3 题（无重复字符的最长子串）
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
*/
func LengthOfLongestSubstring(s string) int {
	// 空串处理
	if s == "" {
		return 0
	}
	// 初始化数据
	window := make(map[byte]int)
	var left, right, res int = 0, 0, 0

	// 滑动窗口
	for right < len(s) {
		c := s[right]
		// 窗口右移
		right++
		// 进行窗口内数据的一系列更新
		window[c]++

		// 判断左侧窗口是否需要收缩
		for window[c] > 1 {
			d := s[left]
			// 窗口左移
			left++
			// 进行窗口内数据的一系列更新
			window[d]--
		}

		// 记录最长无重复子串
		if right-left > res {
			res = right - left
		}
	}

	// 返回结果
	return res
}

/*======================================滑动串口end============================================*/

/*
「力扣」第 1 题（两数之和）
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func twoSum(nums []int, target int) []int {
	// 解法一：之前牛客的题解，使用的是哈希表的思想(不要求原始数组有序)
	var res []int = make([]int, 2, 2)
	// 辅助哈希表
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numCur := nums[i]
		numNext := target - numCur
		if value, ok := hash[numNext]; ok {
			res[0] = value
			res[1] = i
			break
		} else {
			hash[numCur] = i
		}
	}
	return res

	/*
		// 解法二：左右指针，要求原始数组有序
		// 首先对原始数组排序使之有序
			sort.Ints(nums)
			// 然后进行左右指针扫描
			left, right := 0, len(nums)-1
			for left < right {
				sum := nums[left] + nums[right]
				if sum == target {
					res := []int{left, right}
					return res
				} else if sum < target {
					// 让sum大一点
					left++
				} else if sum > target {
					// 让sum小一点
					right--
				}
			}
			// 失败结果（题目表示不会失败，按理走不到这里）
			var tmp []int = []int{-1, -1}
			return tmp
	*/
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
