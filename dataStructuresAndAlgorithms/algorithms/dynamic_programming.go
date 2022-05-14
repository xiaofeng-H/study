package algorithms

/*======================================动态规划 start============================================*/
/*======================================动态规划之背包问题 start============================================*/
/*
	0-1背包
	给你⼀个可装载容量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的容量为 w[i] ，价值为 v[i] ，
   	现在让你⽤这个背包装物品，最多能装的价值是多少？
*/
/*
   	动态规划四步曲：
   	第⼀步：要明确两点，「状态」和「选择」。
			先说状态，如何才能描述一个问题的局面？只要给几个物品和一个背包的容量限制，就形成了一个背包问题呀。
			所以状态就两个，就是 背包的容量 和 可选择的物品。
			再说选择，也很容易想到啊，对于每件物品，你能选择什么？选择就是 装进背包 或者 不装进背包 嘛。
   	第⼆步：要明确 dp 数组的定义。
			首先看看刚才找到的状态，有两个，也就是说我们需要一个二维dp数组。
			dp[i][w]的定义如下：对于前i个物品，当前背包容量为w，这种情况下可以装的的最大价值是dp[i][w]。
			根据这个定义，我们想求的最终答案就是dp[N][W]。base case 就是dp[0][..] = dp[..][0] = 0,因为
			没有物品或者背包没有空间的时候，能装的最大价值就是0。
   	第三步：根据「选择」，思考状态转移的逻辑。
			简单说就是，把物品i装进背包和不把物品i装进背包怎么用代码体现出来呢？
			如果你没有把这第i个物品装进背包，那么很显然，最大价值dp[i][w]应该等于dp[i-1][w]，继承之前的结果。
			如果你把这第i个物品装进了背包，那么dp[i][w]应该等于dp[i-1][w-wt[i-1]] + val[i-1]。而dp[i-1][w-wt[i-1]]
			也很好理解：你如果装了第i个物品，就要寻求剩余重量 w-wt[i-1] 限制下的最大价值，加上第i个物品的价值val[i-1]。
   	最后⼀步：把伪码翻译成代码，处理⼀些边界情况。
*/
func zeroOneKnapsackProblem(W, N int, wt, val []int) int {
	// dp 数组初始化
	dp := make([][]int, N+1, W+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	/*// base case（根据golang语法，int初始值会默认为0，故不需要做这步初始化）
	for i := 0; i <= W; i++ {
		dp[0][i] = 0
	}
	for i := 0; i <= N; i++ {
		dp[i][0] = 0
	}*/
	// 状态转移
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			if (w - wt[i-1]) < 0 {
				// 这种情况下只能选择不装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入背包，择优
				dp[i][w] = maxInt(dp[i-1][w], dp[i-1][w-wt[i-1]]+val[i-1])
			}
		}
	}
	return dp[N][W]
}

/*
「力扣」第 416 题（分割等和子集问题）
时间复杂度：O(N*SUM/2)
空间复杂带：O(N*SUM/2)
*/
// 输⼊⼀个集合，返回是否能够分割成和相等的两个⼦集
func canPartition416(nums []int) bool {
	// 边界值处理
	sum := 0
	for k := range nums {
		sum += nums[k]
	}
	// 和为奇数时，不可能分割成两个和相等的集合
	if sum%2 != 0 {
		return false
	}
	// 状态初始化
	n := len(nums)
	sum /= 2
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
		// base case
		dp[i][0] = true
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if (j - nums[i-1]) < 0 { // 背包容量不足，不能装入第i个物品
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入或者不装入背包
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			}
		}
	}
	return dp[n][sum]
}

/*
「力扣」第 416 题（分割等和子集问题--状态压缩）
时间复杂度：O(N*SUM/2)
空间复杂带：O(SUM)
*/
// 输⼊⼀个集合，返回是否能够分割成和相等的两个⼦集
func canPartition(nums []int) bool {
	// 边界值处理
	sum := 0
	for k := range nums {
		sum += nums[k]
	}
	// 和为奇数时，不可能分割成两个和相等的集合
	if sum%2 != 0 {
		return false
	}
	// 状态初始化
	n := len(nums)
	sum /= 2
	dp := make([]bool, sum+1)
	// base case
	dp[0] = true

	// 状态转移
	for i := 0; i < n; i++ {
		for j := sum; j >= 0; j-- {
			if (j - nums[i]) >= 0 {
				dp[j] = dp[j-nums[i]] || dp[j]
			}
		}
	}
	return dp[sum]
}

/*
	完全背包：没有限制时，一种商品可以无限买，直到背包容量装满
	为什么完全背包和01背包很像？
 	因为01背包在当前商品可以购买时，实际是通过解决自己的i-1的子问题来解决01背包问题，dp[i-1][j - w[i - 1]]，
 	dp[i][j] = Math.max(dp[i - 1][j], dp[i - 1][j - w[i - 1]] + v[i - 1])。
 	完全背包是自己解决自己的i的子问题，dp[i][j - w[i - 1]]。
*/
/*
「力扣」第 518 题（零钱兑换Ⅱ）
时间复杂度：O(N*amount)
空间复杂度：O(N*amount)
*/
/*
	解题思路
	第⼀步要明确两点，「状态」和「选择」。
	状态有两个，就是「背包的容量」和「可选择的物品」，选择就是「装进背包」或者「不装进背包」嘛，背包问题的套路都是这样。
	明⽩了状态和选择，动态规划问题基本上就解决了，只要往这个框架套就完事⼉了：
	for 状态1 in 状态1的所有取值：
		for 状态2 in 状态2的所有取值：
			for ...
				dp[状态1][状态2][...] = 计算(选择1，选择2...)
	第⼆步要明确  dp  数组的定义。
	⾸先看看刚才找到的「状态」，有两个，也就是说我们需要⼀个⼆维  dp数组。
	dp[i][j]  的定义如下：若只使⽤前  i  个物品，当背包容量为  j  时，有  dp[i][j]  种⽅法可以装满背包。
	换句话说，翻译回我们题⽬的意思就是：若只使⽤  coins  中的前  i  个硬币的⾯值，若想凑出⾦额  j  ，有  dp[i][j]  种凑法。
	经过以上的定义，可以得到：
	base case 为  dp[0][..] = 0， dp[..][0] = 1  。因为如果不使⽤任何硬币⾯值，就⽆法凑出任何⾦额；如果凑出的⽬标⾦额为 0，那么“⽆为⽽治”就是
	唯⼀的⼀种凑法。
	我们最终想得到的答案就是  dp[N][amount]  ，其中  N  为  coins  数组的⼤⼩。
	第三步，根据「选择」，思考状态转移的逻辑。
	注意，我们这个问题的特殊点在于物品的数量是⽆限的，所以这⾥和之前写的背包问题⽂章有所不同。
	如果你不把这第  i  个物品装⼊背包，也就是说你不使⽤  coins[i]  这个⾯值的硬币，那么凑出⾯额  j  的⽅法数  dp[i][j]  应该等于  dp[i-1][j]  ，
	继承之前的结果。
	如果你把这第  i  个物品装⼊了背包，也就是说你使⽤  coins[i]  这个⾯值的硬币，那么  dp[i][j]  应该等于  dp[i][j-coins[i-1]]  。
	⾸先由于  i  是从 1 开始的，所以  coins  的索引是  i-1  时表⽰第  i  个硬币的⾯值。dp[i][j-coins[i-1]]  也不难理解，如果你决定使⽤这个⾯值的硬币，那么
	就应该关注如何凑出⾦额  j - coins[i-1]  。
	⽐如说，你想⽤⾯值为 2 的硬币凑出⾦额 5，那么如果你知道了凑出⾦额 3的⽅法，再加上⼀枚⾯额为 2 的硬币，不就可以凑出 5 了嘛。
	综上就是两种选择，⽽我们想求的  dp[i][j]  是「共有多少种凑法」，所以dp[i][j]  的值应该是以上两种选择的结果之和：
	for (int i = 1; i <= n; i++) {
		for (int j = 1; j <= amount; j++) {
			if (j - coins[i-1] >= 0)
				dp[i][j] = dp[i - 1][j] + dp[i][j-coins[i-1]];
	return dp[N][W]
	最后⼀步，把伪码翻译成代码，处理⼀些边界情况。
*/
func completeKnapsackProblem(amount int, coins []int) int {
	n := len(coins) // 面值不同的货币种类数
	// dp 数组初始化
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	/*// base case（根据golang语法，int初始值会默认为0，故不需要做这步初始化）
	for i := 0; i <= amount; i++ {
		dp[0][i] = 0
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}*/
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if (j - coins[i]) >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][amount]
}

//⽽且，我们通过观察可以发现， dp  数组的转移只和  dp[i][..]  和  dp[i- 1][..]  有关，所以可以压缩状态，进⼀步降低算法的空间复杂度：
/*
「力扣」第 518 题（零钱兑换Ⅱ--优化版）
时间复杂度：O(N*amount)
空间复杂度：O(amount)
*/
func change518(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	// base case
	dp[0] = 1
	// 状态转移
	for i := 0; i < n; i++ {
		for j := 1; j <= amount; j++ {
			if (j - coins[i]) >= 0 {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}

/*
   「力扣」第 416 题（分割等和子集问题）
   时间复杂度：O()
*/

/*======================================动态规划之背包问题 end============================================*/

/*
	编辑距离问题就是给我们两个字符串 s1 和 s2 ，只能⽤三种操作，让我们把 s1 变成 s2 ，求最少的操作数。
	注意：解决两个字符串的动态规划问题，一般都是用两个指针i,j分别指向两个字符串的最后，然后一步步往前走，缩小问题规模。
*/
/*
	解题思路

*/
/*
   「力扣」第 72 题（编辑距离）
   时间复杂度：O(M*N)
*/
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1, n+1)
	// base case
	for k := range dp {
		dp[k] = make([]int, n+1)
		dp[k][0] = k
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	// 自底向上求解
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}

		}
	}
	// 储存着整个word1和word2的最小编辑距离
	return dp[m][n]
}
func min(a, b, c int) int {
	return minInt(a, minInt(b, c))
}

/*
「力扣」第 53 题（最大子序和）
*/
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
「力扣」第 322 题（零钱兑换）
*/
func coinChange322(coins []int, amount int) int {
	// dp[i] = x  表⽰，当⽬标⾦额为  i  时，⾄少需要  x  枚硬币
	// 数组大小为amount+1，初始值也为amount+1
	/*PS：为啥 dp 数组初始化为 amount + 1 呢，因为凑成 amount ⾦额的硬币数最多只可能等于 amount （全⽤ 1 元⾯值的硬币），
	  所以初始化为 amount + 1 就相当于初始化为正⽆穷，便于后续取最⼩值。*/
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

/*
「力扣」第 300 题（最长递增子序列）
时间复杂度：O(N*N)
进阶版——时间复杂度：O(NlogN)——patience sorting——了解即可
*/
/*
	最⻓递增⼦序列（Longest Increasing Subsequence，简写 LIS）是⽐较经典的⼀个问题，⽐较容易想到的是动态规划解法，时间复杂度 O(N^2)，
	我们借这个问题来由浅⼊深讲解如何写动态规划。⽐较难想到的是利⽤⼆分查找， 时间复杂度是 O(NlogN)，我们通过⼀种简单的纸牌游戏来辅助理解这种巧妙的解法。
	注意：「⼦序列」和「⼦串」这两个名词的区别，⼦串⼀定是连续的，⽽⼦序列不⼀定是连续的。
*/
func lengthOfLIS300(nums []int) int {
	// 序列长度
	n := len(nums)
	// 定义dp数组，dp[i] 表⽰以 nums[i] 这个数结尾的最⻓递增⼦序列的长度
	dp := make([]int, n)
	// 初始化为1
	for e := range dp {
		dp[e] = 1
	}
	// 结果初始化为0
	res := 0

	// 状态转移
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
	}

	for _, v := range dp {
		if res < v {
			res = v
		}
	}
	return res
}

/*======================================动态规划 end============================================*/
