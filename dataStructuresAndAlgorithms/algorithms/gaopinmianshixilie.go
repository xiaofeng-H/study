package algorithms

import (
	"fmt"
	"math"
	. "study/dataStructuresAndAlgorithms/dataStructures"
)

/*
如何高效地计算素数？
算法名称：Sieve of Eratosthenes(埃拉托色尼筛选法)算法
算法思想：素数的倍数一定不是素数
时间复杂度：O(N*loglogN)
*/
func CountPrimes(n int) int {
	// 数据初始化
	var isPrime = make([]bool, n)
	for k := range isPrime {
		isPrime[k] = true
	}

	// 主算法开始
	for i := 2; i*i < n; i++ {
		// 剪枝：素数的倍数一定不是素数，如果当前判别的数是素数，则说明该数还未筛选（本算法思想是从素数出发的）。
		// 如果当前的数已知不是素数，则说明该数已被筛选过了（默认都是素数），其倍数亦必为该数某个因数的倍数且已被筛选，故无需再次筛选。
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				// j从i平方开始（此处是和第一个for循环的条件前呼后应的，都是为了降低时间复杂度），依次标记i的n倍（+i）不是素数
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
			// 打印素数
			fmt.Printf("%d、", i)
		}
	}
	// 换行
	fmt.Println()

	return count
}

/*
快速模幂算法（力扣372）
算法名称：快速模幂算法
算法思想：模乘法规则---A^2 mod C = (A * A) mod C = ((A mod C) * (A mod C)) mod C
参考链接：https://zh.khanacademy.org/computing/computer-science/cryptography/modarithmetic/a/fast-modular-exponentiation
时间复杂度：O()
*/
const base int = 1337

// 计算a的k次方然后与base求模的结果
func myPow(a, k int) int {
	if k == 0 {
		return 1
	}

	a %= base
	if k%2 == 1 {
		// k是奇数
		return (a * myPow(a, k-1)) % base
	} else {
		// k是偶数
		sub := myPow(a, k/2)
		return (sub * sub) % base
	}
}

func SuperPow372(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]

	part1 := myPow(a, last)
	part2 := myPow(SuperPow372(a, b), 10)
	// 每次乘法都要求模
	return (part1 * part2) % base
}

/*
爱吃香蕉的珂珂（力扣875）
算法名称：
算法思想：⼆分搜索「剪枝」
时间复杂度：O(N*logN)
*/
// 取数组的最大值
func getMax(arr []int) int {
	maxValue := arr[0]
	for _, v := range arr {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

// 时间复杂度：O(N)
func canEat(piles []int, h, speed int) bool {
	time := 0
	for _, v := range piles {
		time += timeOf(v, speed)
	}
	return time <= h
}
func timeOf(n, speed int) int {
	var tmp int = 0
	if n%speed > 0 {
		tmp = 1
	} else {
		tmp = 0
	}
	return (n / speed) + tmp
}
func minEatingSpeed875(piles []int, h int) int {
	// 二分搜索剪枝（套用搜索左侧边界的算法框架）
	left, right := 1, getMax(piles)+1
	for left < right {
		// 防止溢出
		mid := left + (right-left)/2
		if canEat(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

/*
在 D 天内送达包裹的能力（力扣1011）
算法名称：
算法思想：⼆分搜索「剪枝」
时间复杂度：O(N*logN)
*/
// 求一个数组的和
func getSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 如果载重为cap,是否能在D天内运完货物
func canFinish(weight []int, days, cap int) bool {
	i := 0
	for day := 0; day < days; day++ {
		maxCap := cap
		for maxCap -= weight[i]; maxCap >= 0; maxCap -= weight[i] {
			i++
			if i == len(weight) {
				return true
			}
		}
	}
	return false
}

// 寻找左侧边界的二分查找
func shipWithinDays1011(weights []int, days int) int {
	// 载重可能的最小值
	left := getMax(weights)
	// 载重可能的最大值 + 1
	right := getSum(weights) + 1
	// 二分查找剪枝
	for left < right {
		mid := left + (right-left)/2
		if canFinish(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

/*
接雨水（力扣42）--备忘录版本
算法思想：利用备忘录避免重复计算
时间复杂度：O(N)
空间复杂度：O(N)
*/
// 备忘录版本
func trapRemember42(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n := len(height)
	ans := 0
	// 数组充当备忘录
	lMax := make([]int, n)
	rMax := make([]int, n)

	// 从左向右计算lMax
	for i := 1; i < n; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}
	// 从右向左计算rMax
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}
	// 计算答案
	for i := 1; i < n-1; i++ {
		ans += min(lMax[i], rMax[i]) - height[i]
	}
	return ans
}

/*
接雨水（力扣42）--双指针版本
算法思想：牛逼class
时间复杂度：O(N)
空间复杂度：O(1)
*/
// 双指针版本
func trapDoublePoint42(height []int) int {
	if len(height) == 0 {
		return 0
	}

	// 数据初始化
	n := len(height)      // 切片长度
	left, right := 0, n-1 // 双指针
	ans := 0              // 所求结果
	lMax := height[0]     // height[0..left]中最高柱子的高度，即当前柱子左边最高的柱子高度
	rMax := height[n-1]   // height[right..end]的最高柱子的高度，即当前柱子右边最高的柱子高度

	// 开始求解
	for left <= right {
		lMax = max(lMax, height[left])
		rMax = max(rMax, height[right])

		if lMax < rMax {
			ans += lMax - height[left]
			left++
		} else {
			ans += rMax - height[right]
			right--
		}
	}
	return ans
}

/*
删除有序数组中的重复项（力扣26）
算法思想：快慢指针
时间复杂度：O(N)
空间复杂度：O(1)
*/
func removeDuplicates26(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 变量初始化
	n := len(nums)
	slow, fast := 0, 1 // 快慢指针

	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			// 维护nums[0..slow]无重复
			nums[slow] = nums[fast]
		}
		fast++
	}
	// 长度为索引+1
	return slow + 1
}

/*
回文链表（力扣234）
算法思想：借助⼆叉树后序遍历的思路，不需要显式反转原始链表也可以倒序遍历链表，即链表的后序遍历。不过使用该方法的空间复杂度和
反转链表的空间复杂度是一样的，都为O(N)。使用快慢指针可优化空间复杂度，详解看代码。
时间复杂度：O(N)
空间复杂度：O(1)
*/
func isPalindrome234(head *ListNode) bool {
	// 1.先通过双指针技巧中的快慢指针来找到链表的中点；
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow指针现在指向链表的中点（可自己画图模拟该过程来确定slow要不要后移）
	// 2.如果fast指针没有指向nil，说明链表长度为奇数，slow还要再向前进一步；
	if fast != nil {
		slow = slow.Next
	}
	// 3.从slow开始反转后面的链表，现在就可以开始比较回文串了；
	left, right := head, reverseList(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// 头插法
	var cur, pre *ListNode
	cur, pre = head, nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/*
最长回文子串（力扣5）
算法思想：双指针
时间复杂度：O(N^2)
空间复杂度：O(1)
*/
// 寻找以l,r为中心的最长回文子串
func palindrome(s string, l, r int) string {
	// 防止索引越界
	for l >= 0 && r < len(s) && s[l] == s[r] {
		// 向两边展开
		l--
		r++
	}
	// 返回以s[l]和s[r]为中心的最长回文串
	return s[l+1 : r]
}
func longestPalindrome5(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		// 以s[i]为中心的最长回文子串
		s1 := palindrome(s, i, i)
		// 以s[i]和s[i+1]为中心的最长回文子串
		s2 := palindrome(s, i, i+1)
		// res = longest(res,s1,s2)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

/*====================================== 贪心算法 start ============================================*/
// 跳跃游戏参考链接：https://blog.csdn.net/qq_45069496/article/details/123321944
/*
跳跃游戏Ⅰ：是否可以从开始位置跳跃到末尾位置（力扣55）
算法思想：贪心
>>>设想一下，对于数组中的任意一个位置 y，我们如何判断它是否可以到达？根据题目的描述，只要存在一个位置 x，它本身可以到达，
并且它跳跃的最大长度为 x + nums，这个值大于等于 y，即 x + nums[x] ≥ y，那么位置 y 也可以到达。换句话说，
对于每一个可以到达的位置 x，它使得 x + 1 , x + 2 , ⋯ , x + nums[x] 这些连续的位置都可以到达。这样一来，
我们依次遍历数组中的每一个位置，并实时维护最远可以到达的位置。对于当前遍历到的位置 x，如果它在最远可以到达的位置的范围内，
那么我们就可以从起点通过若干次跳跃到达该位置，因此我们可以用 x + nums[x] 更新最远可以到达的位置。在遍历的过程中，
如果最远可以到达的位置大于等于数组中的最后一个位置，那就说明最后一个位置可达，我们就可以直接返回 True 作为答案。反之，
如果在遍历结束后，最后一个位置仍然不可达，我们就返回 False 作为答案。
时间复杂度：O(N)
空间复杂度：O(1)
*/
func canJump55(nums []int) bool {
	// 贪心算法
	// cover表示覆盖下标
	cover := 0
	if len(nums) == 1 {
		return true
	}
	// 从0开始
	i := 0
	for i <= cover { // 卧槽，这个边界值处理牛逼呀（利用for循环的条件判断来处理实际问题中的逻辑判断）（2023/11/4 18:07）
		// cover每次应该在原cover和当前的i+nums[i]取大值，表示覆盖下标
		cover = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		}(cover, i+nums[i])
		// 如果覆盖下标能走到最后一个元素，就可以返回true了
		if cover >= len(nums)-1 {
			return true
		}
		// 更新下标
		i++
	}
	// 下标都走完了，cover依然不能走到最后一个元素，肯定false
	return false
}

/*
跳跃游戏Ⅱ：从开始位置跳到末尾位置最少需要跳多少次（默认可达）（力扣45）
算法思想：贪心
>>>如果我们「贪心」地进行正向查找，每次找到可到达的最远位置，就可以在线性时间内得到最少的跳跃次数。例如，对于数组[2, 3, 1, 2, 4, 2, 3]，
初始位置是下标 0，从下标 0 出发，最远可到达下标 2。下标 0 可到达的位置中，下标 1 的值是 3，从下标 1 出发可以达到更远的位置，
因此第一步到达下标 1.从下标 1 出发，最远可到达下标 4。下标 1 可到达的位置中，下标 4 的值是 4，从下标 4 出发可以达到更远的位置，
因此第二步到达下标4。
>>>在具体的实现中，我们维护当前能够到达的最大下标位置，记为边界。我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
在遍历数组时，我们不访问最后一个元素，这是因为在访问最后一个元素之前，我们的边界一定大于等于最后一个位置，否则就无法跳到最后一个位置了。
如果访问最后一个元素，在边界正好为最后一个位置的情况下，我们会增加一次「不必要的跳跃次数」，因此我们不必访问最后一个元素。
时间复杂度：O(N)
空间复杂度：O(1)
*/
func jump45(nums []int) int {
	n := len(nums)
	// end：上一跳最远可达的位置
	// farthest 上一跳可达的位置进行下一跳可达的最远位置
	// jump：到达当前位置最少的跳跃次数
	end, farthest, jump := 0, 0, 0

	/* 下面代码看起来是一个单循环，但实际表达的思想却更像一个双循环（容易理解）（2023/11/6 15:26）：
	首先在整个区间遍历，找出该区间每一个位置的最远可达距离，而在这个遍历的过程中，其实暗含了一个小区间的划分，那就是当前位置和当前可到达
	的最远位置，在这个小区间中的任意一个位置，是可以由上一跳到达的，所以在当前位置和当前可达的最远位置重合时，就得进行下一次跳跃，
	在选择跳跃的距离时，我们贪心地选择当前这个由上一跳可达的位置中跳的最远的位置。
	*/
	for i := 0; i < n-1; i++ {
		// 每次 for 循环都计算最大可达距离，其实是在方便到达边界时迅速知道下一跳可达的最大距离，
		// 免得到达当前跳跃的边界之后，再去遍历上一跳可达位置上最大可达距离，有点动规的意思了。
		farthest = func(a, b int) int {
			if a > b {
				return a
			} else {
				return b
			}
		}(farthest, nums[i]+i)
		// 到达边界时需要进行下一次跳跃以到达更远的位置
		if end == i {
			jump++
			end = farthest
		}
	}
	return jump
}

/*====================================== 贪心算法 end ============================================*/

/*
k个一组反转链表（力扣25）
算法思想：递归（好好体会吧!2023/10/11 20:50)
时间复杂度：O(n)
空间复杂度：O(1)
*/
// 反转区间[a,b）的元素，注意是左闭右开
func reverse(a, b *ListNode) *ListNode {
	// 头插法
	var pre *ListNode = nil
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	// 返回反转后的头结点
	return pre
}

// 我只能说：递归牛逼（这写法我一辈子都学不会）（2024/5/14 11:24）
func reverseKGroup25(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	// 区间[a,b)包含k个待反转的元素
	a, b := head, head
	for i := 0; i < k; i++ {
		// 不足k个，不需要反转，base case
		if b == nil {
			return head
		}
		b = b.Next
	}

	// 反转前k个元素
	newHead := reverse(a, b)
	// 递归反转后续链表并连接起来
	// 反转之后a到末尾，b的前驱结点到开头，即newHead。下次反转从b开始。
	a.Next = reverseKGroup25(b, k)
	return newHead
}

/* 自己瞎鸡儿实现的代码，简直就是一坨屎，还累个半死，法克！（2024/4/23 17:58）*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	/* 题解：
	   都见了39次了，直接默写！
	*/

	// 边界值处理
	if k == 1 {
		return head
	}

	// 反转开始
	a := head
	b := a
	head = nil
	var cur, tail *ListNode
	i := 1
	for b != nil {
		cur = b.Next
		if i%k == 0 {
			reverseA2B(a, b)
			if head == nil {
				head = b
			} else {
				tail.Next = b
			}
			tail = a
			tail.Next = cur
			a = cur
		}
		i++
		b = cur
	}
	return head
}

// 反转以A开始，以B结尾的字符串
func reverseA2B(a, b *ListNode) {
	// 头插法反转链表
	tail := a     // 反转后的尾结点
	cur := a.Next // 待插入的结点
	a.Next = nil
	var next *ListNode
	for cur != b {
		next = cur.Next
		cur.Next = a
		a = cur
		cur = next
	}
	cur.Next = a
	a = tail
}

/*
有效的括号（力扣20）
算法思想：栈
时间复杂度：
空间复杂度：
*/
// 判断一种括号是否合法
func isValidSingle(s string) bool {
	// 待匹配的左括号数量
	left := 0
	for _, v := range s {
		if v == '(' {
			left++
		} else { // 遇到右括号
			left--
		}

		if left < 0 {
			return false
		}
	}
	return left == 0
}

func isValid20(s string) bool {
	// 辅助栈
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else { // v 是右括号
			length := len(stack)
			if length != 0 {
				c := stackOf(stack)
				if c == s[i] {
					// 出栈
					stack = stack[:length-1]
				} else {
					return false
				}
			} else { // 和最近的左括号不匹配
				return false
			}
		}
	}
	// 是否所有的左括号都匹配了
	return len(stack) == 0
}

func stackOf(stack []byte) byte {
	c := stack[len(stack)-1]
	if c == '(' {
		return ')'
	}
	if c == '[' {
		return ']'
	}
	return '}'
}

/*
消失的数字（力扣面试题17.04）
算法思想：异或运算（哈希解法很简单了）

	异或运算规则：⼀个数和它本⾝做异或运算结果为 0，⼀个数和 0 做异或运算还是它本⾝，
				即a ^ a = 0, a ^ 0 = a，且满足交换律和结合律。

时间复杂度：O(N)
空间复杂度：O(1)
*/
func missingNumber(nums []int) int {
	n := len(nums)
	res := 0

	// 先和新补的索引异或一下
	res ^= n
	for k, v := range nums {
		res ^= k ^ v
	}
	return res
}

/*
错误的集合（力扣645）
算法思想：可使用哈希思想，很简单，不过时间复杂度和空间复杂度都为O(N)。现在使用巧妙解法使得空间复杂度为O(1)。

	通过将每个索引对应的元素变成负数，以表⽰这个索引被对应过⼀次了

时间复杂度：O(N)
空间复杂度：O(1)
*/
func findErrorNums645(nums []int) []int {
	n := len(nums)
	dup := -1
	missing := -1

	for i := 0; i < n; i++ {
		index := int(math.Abs(float64(nums[i]))) - 1
		// nums[index]小于0则说明重复访问
		if nums[index] < 0 {
			dup = index + 1
		} else {
			nums[index] *= -1
		}
	}

	for i := 0; i < n; i++ {
		// nums[i]大于0则说明没有访问
		if nums[i] > 0 {
			// 将索引转换为元素
			missing = i + 1
		}
	}
	return []int{dup, missing}
}

/*
判断子序列（力扣392）
算法思想：双指针
时间复杂度：O(N)
空间复杂度：O(1)
*/
func isSubsequence392(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

/*
判断子序列进阶版本
算法思想：二分
时间复杂度：O(MlogN)
空间复杂度：O(1)
*/
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)
	// 对t进行预处理
	index := make(map[byte][]int)
	// 先记下t中每个字符出现的位置
	for i := 0; i < n; i++ {
		c := t[i]
		if len(index[c]) == 0 {
			index[c] = make([]int, 0)
		}
		index[c] = append(index[c], i)
	}

	// 串t上的指针
	j := 0
	// 借助index查找s[i]
	for i := 0; i < m; i++ {
		c := s[i]
		// 整个t压根没有字符c
		if _, err := index[c]; !err {
			return false
		}
		pos := leftBound(index[c], j)
		// 二分搜索区间中没有找到字符c
		if pos == len(index[c]) {
			return false
		}
		// 向前移动指针j
		j = index[c][pos] + 1
	}
	return true
}

// 查找左侧边界的二分查找
func leftBound(arr []int, tar int) int {
	lo, hi := 0, len(arr)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if tar > arr[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

/*
链表随机结点（力扣382）--未完成(2022/5/9 17:15)
算法思想：栈
时间复杂度：
空间复杂度：
*/
/*type Solution struct {
}

func Constructor(head *ListNode) Solution {

}

// 返回链表中一个随机结点的值
func (this *Solution) GetRandom() int {
	i, res := 0, 0
	p := this

	// for循环遍历链表
	for p != nil {
		// 生成一个[0,i)之间的整数
		// 这个整数等于 0 的概率就是 1/i
		rand.Seed(time.Now().Unix())
		i++
		if rand.Intn(i) == 0 {
			res = p
		}
		p = p
	}
	return res
}*/

/*====================================== Union-Find 算法 start ============================================*/
/*
被围绕的区域（力扣130）
算法思想：常规解法是 DFS（虽然我不会），此处使用并查集即：只有和边界 O 相连的 O 才具有和 dummy 的连通性，他们不会被替换。
时间复杂度：
空间复杂度：
*/
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := len(board), len(board[0])
	// 给dummy留一个额外的位置
	uf := NewUF(m*n + 1)
	dummy := m * n
	// 将首列和末列的0与dummy连通
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(i*n, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.Union(i*n+n-1, dummy)
		}
	}
	// 将首行和末行的0与dummy连通
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			uf.Union(j, dummy)
		}
		if board[m-1][j] == 'O' {
			uf.Union(n*(m-1)+j, dummy)
		}
	}
	// 方向数组d是上下左右搜索的常用手法
	d := [][]int{{1, 0}, {0, 1}, {0, -1}, {-1, 0}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				// 将此0与上下左右的0连通
				for k := 0; k < 4; k++ {
					x := i + d[k][0]
					y := j + d[k][1]
					if board[x][y] == 'O' {
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}
	// 所有不和dummy连通的0，都要被替换
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.Connected(dummy, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}

/*
等式方程的可满足性（力扣990）
算法思想：并查集
时间复杂度：
空间复杂度：
*/
func equationsPossible990(equations []string) bool {
	// 26 个英文字母
	uf := NewUF(26)
	// 先让相等的字母形成连通分量
	for _, v := range equations {
		if v[1] == '=' {
			x := v[0] - 'a'
			y := v[3] - 'a'
			uf.Union(int(x), int(y))
		}
	}
	// 检查不等关系是否打破相等关系的连通性
	for _, v := range equations {
		if v[1] == '!' {
			x := v[0] - 'a'
			y := v[3] - 'a'
			// 如果相等关系成立，就是逻辑冲突
			if uf.Connected(int(x), int(y)) {
				return false
			}
		}
	}
	return true
}

/*====================================== Union-Find 算法 end ============================================*/

/*====================================== 脑筋急转弯 start ============================================*/
/*
Nim游戏（力扣292）
算法思想：我们解决这种问题的思路⼀般都是反着思考：
如果我能赢，那么最后轮到我取⽯⼦的时候必须要剩下 1~3 颗⽯⼦，这样我才能⼀把拿完。
如何营造这样的⼀个局⾯呢？显然，如果对⼿拿的时候只剩 4 颗⽯⼦，那么⽆论他怎么拿，总会剩下 1~3 颗⽯⼦，我就能赢。
如何逼迫对⼿⾯对 4 颗⽯⼦呢？要想办法，让我选择的时候还有 5~7 颗⽯⼦，这样的话我就有把握让对⽅不得不⾯对 4 颗⽯⼦。
如何营造 5~7 颗⽯⼦的局⾯呢？让对⼿⾯对 8 颗⽯⼦，⽆论他怎么拿，都会给我剩下 5~7 颗，我就能赢。
这样⼀直循环下去，我们发现只要踩到 4 的倍数，就落⼊了圈套，永远逃不出 4 的倍数，⽽且⼀定会输。所以这道题的解法⾮常简单：
时间复杂度：O(1)
空间复杂度：O(1)
*/
func canWinNim292(n int) bool {
	// 如果上来就踩到4的倍数，那就认输吧
	// 否则，可以把对方控制在4的倍数，必胜
	return n%4 != 0
}

/*
石子游戏（力扣877）
算法思想：只要你⾜够聪明，你是必胜⽆疑的，因为你是先⼿。
这是为什么呢，因为题⽬有两个条件很重要：⼀是⽯头总共有偶数堆，⽯头的总数是奇数。这两个看似增加游戏公平性的条件，反⽽使该游戏成为了⼀
个割⾲菜游戏。我们以  piles=[2, 1, 9, 5]  讲解，假设这四堆⽯头从左到右的索引分别是 1，2，3，4。如果我们把这四堆⽯头按索引的奇偶
分为两组，即第 1、3 堆和第 2、4 堆，那么这两组⽯头的数量⼀定不同，也就是说⼀堆多⼀堆少。因为⽯头的总数是奇数，不能被平分。⽽作为
第⼀个拿⽯头的⼈，你可以控制⾃⼰拿到所有偶数堆，或者所有的奇数堆。你最开始可以选择第 1 堆或第 4 堆。如果你想要偶数堆，你就拿第 4 堆，
这样留给对⼿的选择只有第 1、3 堆，他不管怎么拿，第 2 堆⼜会暴露出来，你就可以拿。同理，如果你想拿奇数堆，你就拿第 1 堆，留给对⼿的只有第
2、4 堆，他不管怎么拿，第 3 堆⼜给你暴露出来了。也就是说，你可以在第⼀步就观察好，奇数堆的⽯头总数多，还是偶数堆的⽯头总数多，然后步步为营，
就⼀切尽在掌控之中了。知道了这个漏洞，可以整⼀整不知情的同学了。
时间复杂度：O(1)
空间复杂度：O(1)
*/
func stoneGame(piles []int) bool {
	return true
}

/*
灯泡开关（力扣319）
算法思想：什么？这个问题跟平⽅根有什么关系？其实这个解法挺精妙，如果没⼈告诉你解法，还真不好想明⽩。
⾸先，因为电灯⼀开始都是关闭的，所以某⼀盏灯最后如果是点亮的，必然要被按奇数次开关。我们假设只有 6 盏灯，⽽且我们只看第 6 盏灯。
需要进⾏ 6 轮操作对吧，请问对于第 6 盏灯，会被按下⼏次开关呢？这不难得出，第 1 轮会被按，第 2 轮，第 3 轮，第 6 轮都会被按。
为什么第 1、2、3、6 轮会被按呢？因为  6=1*6=2*3  。⼀般情况下，因⼦都是成对出现的，也就是说开关被按的次数⼀般是偶数次。但是有特殊情况，
⽐如说总共有 16 盏灯，那么第 16 盏灯会被按⼏次? 16=1*16=2*8=4*4，其中因⼦ 4 重复出现，所以第 16 盏灯会被按 5 次，奇数次。现在你应该理
解这个问题为什么和平⽅根有关了吧？不过，我们不是要算最后有⼏盏灯亮着吗，这样直接平⽅根⼀下是啥意思呢？稍微思考⼀下就能理解了。
就假设现在总共有 16 盏灯，我们求 16 的平⽅根，等于 4，这就说明最后会有 4 盏灯亮着，它们分别是第 1*1=1 盏、第 2*2=4 盏、第 3*3=9 盏和第
4*4=16 盏。就算有的 n 平⽅根结果是⼩数，强转成 int 型，也相当于⼀个最⼤整数上界，⽐这个上界⼩的所有整数，平⽅后的索引都是最后亮着的灯的索引。
所以说我们直接把平⽅根转成整数，就是这个问题的答案。
时间复杂度：O(1)
空间复杂度：O(1)
*/
func bulbSwitch319(n int) int {
	return int(math.Sqrt(float64(n)))
}

/*====================================== 脑筋急转弯 end ============================================*/
