package algorithms

import (
	"fmt"
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
「力扣」第 206 题（反转链表）
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList206(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 头插法
	var q *ListNode = nil
	for head != nil {
		p := head.Next
		head.Next = q
		q = head
		head = p
	}
	return q
}

/*
「力扣」第 344 题（反转字符串）
*/
func reverseString344(s []byte) {
	if len(s) == 0 {
		fmt.Println("")
	}

	// 双指针
	left, right := 0, len(s)-1
	for left <= right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
	fmt.Println(s)
}

/*
「力扣」第 88 题（合并两个有序数组）
*/
func merge88(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	// 初始化
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		arr[i] = nums1[i]
	}
	i, j, k := 0, 0, 0

	// 归并
	for i < m && j < n {
		if arr[i] <= nums2[j] {
			nums1[k] = arr[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	// nums1还未归并完
	for i < m {
		nums1[k] = arr[i]
		i++
		k++
	}
	// nums2还未归并完
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

/*
「力扣」第 21 题（合并两个有序链表）
*/
func mergeTwoLists21(list1 *ListNode, list2 *ListNode) *ListNode {
	// 边界值处理
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}

	var head = new(ListNode)
	var pre = head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			pre = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			pre = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return head.Next
}

/*======================================二叉树 start============================================*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
「力扣」第 94 题（二叉树中序遍历）
*/
func inorderTraversal94(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 辅助栈
	stack := make([]*TreeNode, 0)
	// 中序遍历序列
	res := make([]int, 0)
	// 当前被操作结点
	element := root

	// 当前被操作结点或者辅助栈不为空时进行中序遍历
	for element != nil || len(stack) != 0 {
		// 当前结点左子树入栈
		for element != nil {
			stack = append(stack, element)
			element = element.Left
		}
		// 当前辅助栈长度
		length := len(stack)
		// 访问栈顶结点
		if length != 0 {
			// 栈顶结点出栈
			p := stack[length-1]
			stack = stack[:length-1]
			res = append(res, p.Val)
			// 若出栈元素的右孩子不空，则将当前被操作元素指向该右孩子
			if p.Right != nil {
				element = p.Right
			}
		}
	}
	return res
}

/*
「力扣」第 144 题（二叉树前序遍历）
*/
func preorderTraversal144(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 辅助栈
	stack := make([]*TreeNode, 0)
	// 前序遍历序列
	res := make([]int, 0)
	// 根结点入栈
	stack = append(stack, root)

	// 栈不空时进行前序遍历
	for len(stack) != 0 {
		length := len(stack)
		// 栈顶结点出栈
		p := stack[length-1]
		stack = stack[:length-1]
		res = append(res, p.Val)
		// 右孩子先入栈，左孩子再入栈（先进后出）
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return res
}

/*
「力扣」第 145 题（二叉树后序遍历）
*/
func postorderTraversal145(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	// 辅助栈
	stack1 := make([]*TreeNode, 0)
	stack2 := make([]int, 0)
	// 前序遍历序列
	res := make([]int, 0)
	// 根结点入栈
	stack1 = append(stack1, root)

	// 栈不空时进行后序遍历
	for len(stack1) != 0 {
		length := len(stack1)
		// 栈顶结点出栈
		p := stack1[length-1]
		stack1 = stack1[:length-1]
		stack2 = append(stack2, p.Val)
		// 左孩子入栈，右孩子再入栈
		if p.Left != nil {
			stack1 = append(stack1, p.Left)
		}
		if p.Right != nil {
			stack1 = append(stack1, p.Right)
		}
	}

	// stack2中保存的是逆后序序列
	for len(stack2) != 0 {
		length := len(stack2)
		res = append(res, stack2[length-1])
		stack2 = stack2[:length-1]
	}
	return res
}

/*
「力扣」第 102 题（二叉树层次遍历）
*/
func levelOrder102(root *TreeNode) [][]int {
	// 边界值处理
	if root == nil {
		return nil
	}

	// 辅助变量初始化
	queue := make([]*TreeNode, 0) // 辅助队列
	res := make([][]int, 0)       // 遍历结果

	// 根结点入队
	queue = append(queue, root)
	// 队不空时进行层次遍历
	for len(queue) > 0 {
		// 同一层的结点出队
		arr := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]
			arr = append(arr, node.Val)
			// 该结点的左右孩子入队
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, arr)
	}
	return res
}

/*
「力扣」第 98 题（验证二叉搜索树）
*/
func isValidBST98(root *TreeNode) bool {
	return isValidBSTPlus(root, nil, nil)
}

func isValidBSTPlus(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBSTPlus(root.Left, min, root) && isValidBSTPlus(root.Right, root, max)
}

/*
「力扣」第 100 题（相同的树）
*/
func isSameTree100(p *TreeNode, q *TreeNode) bool {
	// 边界值处理（注意以下两个if条件的并列使用）
	// 都为空的话，显然相同
	if p == nil && q == nil {
		return true
	}
	// 一个为空，一个非空，显然不同
	if p == nil || q == nil {
		return false
	}
	// 两个都非空，但val不一样也不行
	if p.Val != q.Val {
		return false
	}
	// p和q该比的都比完了
	return isSameTree100(p.Left, q.Left) && isSameTree100(p.Right, q.Right)
}

/*
「力扣」第 700 题（二叉搜索树中的搜索）
*/
func searchBST700(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	// 剪枝（二叉搜索树性质）
	if root.Val < val {
		return searchBST700(root.Right, val)
	}
	return searchBST700(root.Left, val)
	// root该做的事做完了，顺带把框架也完成了，妙
}

/*
「力扣」第 701 题（二叉搜索树中的插入操作）
*/
func insertIntoBST701(root *TreeNode, val int) *TreeNode {
	// 找到空位置插入新结点
	if root == nil {
		return &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
	}
	// if root.Val == val // BST中一般不会插入已存在的元素
	// 二叉搜索树的剪枝操作
	if root.Val < val {
		root.Right = insertIntoBST701(root.Right, val)
	}
	if root.Val > val {
		root.Left = insertIntoBST701(root.Left, val)
	}
	return root
}

/*
「力扣」第 450 题（二叉搜索树中的删除操作）
*/
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == key {
		// 这两个if把情况1和2都正确处理了（这逻辑，牛逼）2022/5/12 21:53
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// 处理情况3
		minNode := getMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, minNode.Val)
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func getMin(node *TreeNode) *TreeNode {
	// BST最左边的就是最小值
	for node.Left != nil {
		node = node.Left
	}
	return node
}

/*
「力扣」第 222 题（完全二叉树的结点个数）
时间复杂度：O(logN*logN)
*/
func countNodes222(root *TreeNode) int {
	l, r := root, root
	// 记录左右子树的高度
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	// 如果左右子树的高度相同，则是一颗满二叉树
	if hl == hr {
		return int(math.Pow(2, float64(hl)) - 1)
	}
	// 如果左右高度不同，则按照普通二叉树的逻辑计算
	return 1 + countNodes222(root.Left) + countNodes222(root.Right)
}

/*======================================二叉树 end============================================*/

/*
「力扣」第 27 题（移除元素）
*/
func removeElement27(nums []int, val int) int {
	// 快慢指针
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

/*
「力扣」第 58 题（最后一个单词的长度）
*/
func lengthOfLastWord58(s string) int {
	// 先统计字符串末尾空格数
	counts := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			counts++
		} else {
			break
		}
	}
	length := 0
	for i := 0; i < len(s)-counts; i++ {
		if s[i] != ' ' {
			length++
		} else {
			length = 0
		}
	}
	return length
}

/*
「力扣」第 125 题（验证回文串）
*/
func isPalindrome125(s string) bool {
	// 空串默认为回文串
	if len(s) == 0 {
		return true
	}

	// 先将原始串统一转换为小写
	str := strings.ToLower(s)
	// 采用双指针进行回文串判定
	left, right := 0, len(str)-1
	for left < right {
		// 先检查双指针是否指向字母或者数字
		for left < right && !isNumsOrWords(str[left]) {
			left++
		}
		for left < right && !isNumsOrWords(str[right]) {
			right--
		}
		// 判断是否是回文串
		if left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

// 判断一个字符是否属于数字或者字母
func isNumsOrWords(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

/*
「力扣」第 141 题（环形链表）
*/
func hasCycle141(head *ListNode) bool {
	// 快慢指针
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

/*======================================动态规划 start============================================*/
/*
「力扣」第 322 题（零钱兑换）
*/
func coinChange322(coins []int, amount int) int {
	// dp[i] = x  表⽰，当⽬标⾦额为  i  时，⾄少需要  x  枚硬币
	// 数组大小为amount+1，初始值也为amount+1
	dp := make([]int, amount+1)
	for k := range dp {
		dp[k] = amount + 1
	}

	// base case
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		// 内层for循环在求所有子问题+1的最小值
		for _, v := range coins {
			// 子问题无解，跳过
			if i-v < 0 {
				continue
			}
			dp[i] = minInt(dp[i], 1+dp[i-v])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

/*======================================动态规划 end============================================*/

/*======================================回溯 start============================================*/
/*
「力扣」第 46 题（全排列）
*/
var resPermute [][]int

func Permute46(nums []int) [][]int {
	resPermute = make([][]int, 0)
	// 记录路径
	track := make([]int, 0)
	backTrackPermute(nums, track)
	return resPermute
}

// 路径：记录在track中
// 选择列表：nums中不存在与track的那些元素
// 结束条件：nums中的元素全都在track中出现
func backTrackPermute(nums, track []int) {
	// 触发结束条件
	if len(track) == len(nums) {
		fmt.Printf("满足条件的全排列为：%v\n", track)
		fmt.Printf("添加前的结果集为：%v\n", resPermute)
		// 注意：切片本身为引用类型，修改值会修改底层数组，导致添加到结果集中的数据也发生改变，使得结果集会有重复结果，
		// 实际是因为修改切片值引起的，从而引发bug(找了好半天)！需要新开辟空间，以永久存储结果。（2022/5/12 23:42）
		arr := make([]int, len(track))
		for k := range track {
			arr[k] = track[k]
		}
		resPermute = append(resPermute, arr)
		fmt.Printf("添加后的结果集为：%v\n", resPermute)
		return
	}

	for i := 0; i < len(nums); i++ {
		// 排除不合法的选择
		if containsInt(track, nums[i]) {
			continue
		}
		// 做选择
		track = append(track, nums[i])
		fmt.Printf("选择后的路径为：%v,新增的选择为：%v\n", track, nums[i])
		// 进入下一层决策树
		backTrackPermute(nums, track)
		// 撤销选择
		length := len(track)
		track = track[:length-1]
		fmt.Printf("撤销选择后的路径为：%v,撤销的选择为：%v\n", track, nums[i])
	}
}

// 判断切片中是否存在target
func containsInt(arr []int, target int) bool {
	for k := range arr {
		if arr[k] == target {
			return true
		}
	}
	return false
}

/*
「力扣」第 51 题（N皇后）
*/
var resNQueens [][]string // 储存结果
// 输入棋盘边长 n，返回所有合法的位置
func solveNQueens(n int) [][]string {
	// 经典回溯算法
	// 初始化
	resNQueens = make([][]string, 0)
	// '.' 表示空，'Q' 表示皇后，初始化空棋盘
	board := make([][]byte, n) // 棋盘初始化
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	backTrackNQueens(board, 0)
	return resNQueens
}

// 路径：board中小于row的那些行都已经成功放置了皇后
// 选择列表：第row行的所有列都是放置皇后的选择
// 结束条件：row超过board的最后一行
func backTrackNQueens(board [][]byte, row int) {
	// 结束条件
	if row == len(board) {
		// 记录棋盘分布并保存结果
		str := make([]string, 0)
		for _, v := range board {
			s := string(v)
			str = append(str, s)
		}
		resNQueens = append(resNQueens, str)
		return
	}

	for col := 0; col < len(board); col++ {
		// 排除不合法选择
		if !isValidLeetCode(board, row, col) {
			continue
		}
		// 做出选择
		board[row][col] = 'Q'
		// 进入下一行决策
		backTrackNQueens(board, row+1)
		// 撤销选择
		board[row][col] = '.'
	}
}

// 判断当前位置是否可以继续放置皇后
func isValidLeetCode(board [][]byte, row, col int) bool {
	// 先看当前列是否已经放置皇后
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 再看左上方是否已经放置皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	// 再看右上方是否已经放置皇后
	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}
	return true
}
/*======================================回溯 end============================================*/

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

/*
「力扣」第 2148 题（元素计数）
*/

/*
「力扣」第 2148 题（元素计数）
*//*
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
*//*
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
