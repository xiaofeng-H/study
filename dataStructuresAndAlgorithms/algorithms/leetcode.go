package algorithms

import (
	"math"
	"strconv"
	"strings"
)

/*
「力扣」第 53 题（最大子序和）
*/
// 动态规划
func MaxSubArray53(nums []int) int {
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
func ClimbStairs70(n int) int {
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
func Generate118(numRows int) [][]int {
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
func GetRow119(rowIndex int) []int {
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
func CountElements2148(nums []int) int {
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
func Calculate17(s string) int {
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
func RotateString796(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	ss := s + s
	return strings.Contains(ss, goal)
}

/*
「力扣」第 1108 题（IP地址无效化）
*/
func DefangIPaddr1108(address string) string {
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
func MinWindow76(s string, t string) string {
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
func CheckInclusion567(s1 string, s2 string) bool {
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
func FindAnagrams438(s string, p string) []int {
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
func LengthOfLongestSubstring3(s string) int {
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
func TwoSum1(nums []int, target int) []int {
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

/*======================================BFS start============================================*/
/*
「力扣」第 111 题（二叉树的最小深度）
*/
// 二叉树结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS算法
func MinDepth111(root *TreeNode) int {
	// 空树处理
	if root == nil {
		return 0
	}
	// 变量初始化
	var queue []*TreeNode = make([]*TreeNode, 0)
	var front, rear = -1, -1 // 队列的首尾下标
	var depth int = 1        // root本身就是一层，depth初始化为1
	// BFS核心代码
	// 根结点入队
	rear++
	queue = append(queue, root)
	for rear != front {
		length := rear - front
		// 将当前队列中的所有结点向四周扩散（一层结点出队）
		for i := 0; i < length; i++ {
			// 队头结点出队
			front++
			cur := queue[front]
			// 判断是否到达终点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			// 将cur的相邻结点加入队列
			if cur.Left != nil {
				rear++
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				rear++
				queue = append(queue, cur.Right)
			}
		}
		// 这里增加步数
		depth++
	}
	return depth
}

/*
「力扣」第 752 题（打开转盘锁）
你有一个带有四个圆形拨轮的转盘锁。每个拨轮都有10个数字： '0', '1', '2', '3', '4', '5', '6', '7', '8', '9' 。每个拨轮可以自由旋转：例如把 '9' 变为 '0'，'0' 变为 '9' 。每次旋转都只能旋转一个拨轮的一位数字。
锁的初始数字为 '0000' ，一个代表四个拨轮的数字的字符串。
列表 deadends 包含了一组死亡数字，一旦拨轮的数字和列表里的任何一个元素相同，这个锁将会被永久锁定，无法再被旋转。
字符串 target 代表可以解锁的数字，你需要给出解锁需要的最小旋转次数，如果无论如何不能解锁，返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/open-the-lock
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 将s[j]向上拨动一次
func plusOne(s string, j int) string {
	ch := []byte(s)
	if ch[j] == '9' {
		ch[j] = '0'
	} else {
		ch[j]++
	}
	return string(ch)
}

// 将s[i]向下拨动一次
func minusOne(s string, i int) string {
	ch := []byte(s)
	if ch[i] == '0' {
		ch[i] = '9'
	} else {
		ch[i]--
	}
	return string(ch)
}

// BFS框架，打印出所有可能的密码
// 该法内存溢出，队列处理欠佳
func OpenLock752(deadends []string, target string) int {
	// 记录需要跳过的死亡密码
	dead := make(map[string]int)
	for k, v := range deadends {
		dead[v] = k
	}
	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]int)
	// 初始化数据
	q := make([]string, 0)
	front, rear := -1, -1
	// 从起点开始启动广度优先搜索
	step := 0
	// 初始密码入队并记录
	rear++
	q = append(q, "0000")
	visited["0000"] = 1

	for front != rear {
		length := rear - front
		// 将当前队列中的所有结点向周围扩散
		for i := 0; i < length; i++ {
			// 队头结点出队
			front++
			cur := q[front]
			// 判断是否到达终点
			if _, ok := dead[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			// 将一个结点的未遍历相邻结点加入队列
			for j := 0; j < 4; j++ {
				up := plusOne(cur, j)
				if _, ok := visited[up]; !ok {
					rear++
					q = append(q, up)
				}
				down := minusOne(cur, j)
				if _, ok := visited[down]; !ok {
					rear++
					q = append(q, down)
				}
			}
		}
		// 在这里增加步数
		step++
	}
	// 如果穷举完都没有找到目标密码，那就是找不到了
	return -1
}

func openLock(deadends []string, target string) int {
	step := 0 // 旋转次数
	deadendsMap := make(map[string]bool)
	visitedMap := make(map[string]bool)

	for _, v := range deadends { // 记录所有“死亡点”
		deadendsMap[v] = true
	}

	q := []string{"0000"} // 队列q
	for len(q) > 0 {      // 循环直至队列为空
		size := len(q)              // 获取BFS当前level的节点个数
		for i := 0; i < size; i++ { // 遍历当前层的节点
			node := q[0]        // 获取出列的节点
			q = q[1:]           // 节点出列
			if node == target { // 如果出列的节点正好是目标节点
				return step // 返回当前所用的步数
			}
			if _, ok := visitedMap[node]; ok { // 之前访问过该节点，跳过
				continue
			}
			if _, ok := deadendsMap[node]; ok { // 遇到“死亡点”，跳过
				continue
			}
			visitedMap[node] = true // 将该点标记为访问过

			for j := 0; j < len(node); j++ { // 通过遍历当前字符串，找出它的所有子节点，安排入列
				num := int(node[j] - '0')                             // 获取当前的数字num
				up := (num + 1) % 10                                  // 往上拧所得的新数，比如1变成2
				down := (num + 9) % 10                                // 往下拧所得的新数，比如7变成6
				q = append(q, node[:j]+strconv.Itoa(up)+node[j+1:])   // 拼成新字符串，入列
				q = append(q, node[:j]+strconv.Itoa(down)+node[j+1:]) // 拼成新字符串 入列
			}
		}
		step++ // 当前层的所有节点遍历完毕，层次+1
	}
	return -1 // 无论如何都遇不到目标节点，返回-1
}

// 双向BFS
func OpenLockBothway752(deadends []string, target string) int {
	// 数据初始化
	dead := make(map[string]bool)
	for _, v := range deadends {
		dead[v] = true
	}
	// 用集合不用队列，可以快速判断元素是否存在
	q1 := make(map[string]bool)
	q2 := make(map[string]bool)
	visited := make(map[string]bool)
	step := 0
	q1["0000"] = true
	q2[target] = true

	for len(q1) != 0 && len(q2) != 0 {
		// 哈希表在遍历的过程中不能修改，用tmp存储扩散结果
		tmp := make(map[string]bool)

		// 将q1中的所有结点向四周扩散
		for k := range q1 {
			// 判断是否到达终点
			if _, ok := dead[k]; ok {
				continue
			}
			if _, ok := q2[k]; ok {
				return step
			}
			visited[k] = true

			// 将一个结点的未遍历相邻结点加入集合
			for j := 0; j < 4; j++ {
				up := plusOne(k, j)
				if _, ok := visited[up]; !ok {
					tmp[up] = true
				}
				down := minusOne(k, j)
				if _, ok := visited[down]; !ok {
					tmp[down] = true
				}
			}
		}
		/* 在这⾥增加步数 */
		step++
		// temp 相当于 q1
		// 这⾥交换 q1 q2，下⼀轮 while 就是扩散 q2
		q1 = q2
		q2 = tmp
	}
	return -1
}

/*======================================BFS end============================================*/

/*======================================递归 start============================================*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归初识
// 给他⼀个节点和⼀个⽬标值，他返回以这个节点为根的树 中，和为⽬标值的路径总数
func pathSum437(root *TreeNode, targetSum int) int {
	// 递归结束条件
	if root == nil {
		return 0
	}

	// 自己为开头的路径数
	pathImLeading := count(root, targetSum)
	// 左边路径总数（相信他能算出来）
	leftPathSum := pathSum437(root.Left, targetSum)
	// 右边路径总数（相信他也能算出来）
	rightPathSum := pathSum437(root.Right, targetSum)

	return pathImLeading + leftPathSum + rightPathSum

}

// 给他⼀个节点和⼀个⽬标值，他返回以这个节点为根的树中， 能凑出⼏个以该节点为路径开头，和为⽬标值的路径总数
func count(node *TreeNode, sum int) int {
	// 递归结束条件
	if node == nil {
		return 0
	}

	// 我自己能不能独当一面，作为一条单独的路径呢？
	var isMe int
	if node.Val == sum {
		isMe = 1
	} else {
		isMe = 0
	}
	// 左边的小老弟，你那边能凑出几个 sum - node.Val 呀？
	leftBrother := count(node.Left, sum-node.Val)
	// 右边的小老弟，你那边能凑出几个 sum - node.Val 呀？
	rightBrother := count(node.Right, sum-node.Val)

	// 我这能凑这么多个
	return isMe + leftBrother + rightBrother
}

/*
数组中的逆序对（剑指51）
算法名称：分治
算法思想：二路归并排序
时间复杂度：O(N*logN)
*/
func reversePairs(nums []int) int {
return 0
}
/*======================================递归 end============================================*/

/*======================================分治 start============================================*/
/*======================================分治 end============================================*/

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
