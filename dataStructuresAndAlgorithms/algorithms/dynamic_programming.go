package algorithms

import "fmt"

/*====================================== 动态规划 start ============================================*/
/*
	之前认识了一个女孩，不知不觉已经喜欢五年了。刚开始只是在 QQ 上聊天，说实话也没太多的话题，大多数时间属于尬聊，我问一句人家回一句。
记得是 19 年五一节前吧，有一次聊得比较晚，我也趁机微露了自己的心意，忘了当时她是怎么回应的了，只记得此后聊天的频率骤降，仅在逢节寒暄一两句罢了。
后来是 20 年 1 月 20 刚好认识了一年，我便给她推荐了五首歌（我两都喜欢听歌），是一句藏头诗“一周年快乐”，或许在她看来不过是几首无聊歌曲罢了，
一如往常，礼貌地回了一句。忘记后来发生了什么事，便断了联系。直到 21 年 11 月，偶然间在朋友圈看到她正在集赞免费领取考研资料，恰巧我之前准备过
手头上有一些资料，便发了私信，这应该是后来差不多两年时间第一次联系吧，她依旧很有礼貌向我道谢，此后也曾联系过几次，大抵是一些备考建议，ps:也曾偷偷
给她送了两箱特仑苏。那年冬天疫情还很严重，我一个朋友从天津回家过年，途中经过天水去挽回已经分手将近半年的女友而被隔离（当然这才是他的主要目的，
两人现已鸳鸯成对了），一个人形单影只无依无靠，哈哈哈。。。她是天水人，我便借此机会劳驾她买些日用品送给我朋友，她很爽快的答应了（几桶汤达人、火腿肠、辣片、
果粒橙、哇哈哈、笔记本...），之后转她钱起初不愿收，最后才说六十块（据我朋友说应该不止六十吧）。再后来 22 年 2 月，在我发小的怂恿下去天水找人家，
人家没来，删了所有的联系方式，自此断了音信，到今天 21 个月时间。我也经常问自己为啥总是念念不忘？反正绝不是因为相信张宇老师说的“念念不忘，必有回响”哈！
或许大概是在人生的这段旅途中没遇到其他人吧，亦或许是忘不掉吧。我不知道，我是个怀旧的人，我甚至没有见过她。（2023/11/14 17:03）
*/
/*====================================== 动态规划之背包问题 start ============================================*/
/*
	0-1背包
	给你⼀个可装载容量为 W 的背包和 N 个物品，每个物品有重量和价值两个属性。其中第 i 个物品的容量为 w[i] ，价值为 v[i] ，
   	现在让你⽤这个背包装物品，最多能装的价值是多少？
*/
/*
   	动态规划四步曲：
   	第⼀步：要明确两点，「状态」和「选择」。
			先说状态，如何才能描述一个问题的局面？只要给几个物品和一个背包的容量限制，就形成了一个背包问题呀。
			所以状态就两个，就是 背包的容量 和 可选择的物品。（可理解为此问题的变量）
			再说选择，也很容易想到啊，对于每件物品，你能选择什么？选择就是 装进背包 或者 不装进背包 嘛。
   	第⼆步：要明确 dp 数组的定义。
			首先看看刚才找到的状态，有两个，也就是说我们需要一个二维 dp 数组。
			dp[i][w] 的定义如下：对于前 i 个物品，当前背包容量为 w，这种情况下可以装的的最大价值是 dp[i][w]。
			根据这个定义，我们想求的最终答案就是 dp[N][W]。base case 就是 dp[0][..] = dp[..][0] = 0,因为
			没有物品或者背包没有空间的时候，能装的最大价值就是 0。
   	第三步：根据「选择」，思考状态转移的逻辑。
			简单说就是，把物品 i 装进背包和不把物品 i 装进背包怎么用代码体现出来呢？
			如果你没有把这第 i 个物品装进背包，那么很显然，最大价值 dp[i][w] 应该等于 dp[i-1][w]，继承之前的结果。
			如果你把这第 i 个物品装进了背包，那么 dp[i][w] 应该等于 dp[i-1][w-wt[i-1]] + val[i-1]。而 dp[i-1][w-wt[i-1]]
			也很好理解：你如果装了第 i 个物品，就要寻求剩余重量 w-wt[i-1] 限制下的最大价值，加上第 i 个物品的价值 val[i-1]。
   	最后⼀步：把伪码翻译成代码，处理⼀些边界情况。
*/
func zeroOneKnapsackProblem(W, N int, wt, val []int) int {
	// dp 数组初始化
	// dp[i][j]：表示在前i个物品可选择装入的情况下，当背包容量为j时，可装入的最大价值为dp[i][j]。
	dp := make([][]int, N+1, W+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
	}
	// base case
	for i := 0; i <= W; i++ {
		// 没有物品可装入时，背包的最大价值为0
		dp[0][i] = 0
	}
	for i := 0; i <= N; i++ {
		// 背包容量为0时，不能装入物品，故背包的最大价值为0
		dp[i][0] = 0
	}
	// 状态转移
	for i := 1; i <= N; i++ {
		for w := 1; w <= W; w++ {
			// w：当前背包容量；wt[i-1]：当前准备加入的物品的体积
			if (w - wt[i-1]) < 0 {
				// 1.若当前物品体积大于背包总容量，则不能装入背包
				dp[i][w] = dp[i-1][w]
			} else {
				// 2.否则，择优选择装入或者不装入背包
				dp[i][w] = MaxIntAB(dp[i-1][w], dp[i-1][w-wt[i-1]]+val[i-1])
			}
		}
	}
	// dp[N][W]：表示背包容量为W且可选择N个物品时，背包能装下的最大价值。
	return dp[N][W]
}

/*
	完全背包：没有限制时，一种商品可以无限买，直到背包容量装满
	为什么完全背包和01背包很像？
 	因为01背包在当前商品可以购买时，实际是通过解决自己的 i-1 的子问题来解决01背包问题，dp[i-1][j - w[i - 1]]，
 	dp[i][j] = Math.max(dp[i - 1][j], dp[i - 1][j - w[i - 1]] + v[i - 1])。
 	完全背包是自己解决自己的 i 的子问题，dp[i][j - w[i - 1]]。
*/
/*
「力扣」第 518 题（零钱兑换Ⅱ）
给你一个整数数组 coins 表示不同面额的硬币，另给一个整数 amount 表示总金额。请你计算并返回可以凑成总金额的硬币组合数。
如果任何硬币组合都无法凑出总金额，返回 0 。
假设每一种面额的硬币有无限个。
时间复杂度：O(N*amount)
空间复杂度：O(N*amount)
*/
/*
	解题思路
	第⼀步:要明确两点，「状态」和「选择」。
	状态有两个，就是「背包的容量」和「可选择的物品」，选择就是「装进背包」或者「不装进背包」嘛，背包问题的套路都是这样。
	明⽩了状态和选择，动态规划问题基本上就解决了，只要往这个框架套就完事⼉了：
	for 状态1 in 状态1的所有取值：
		for 状态2 in 状态2的所有取值：
			for ...
				dp[状态1][状态2][...] = 计算(选择1，选择2...)
	第⼆步:要明确 dp 数组的定义。
	⾸先看看刚才找到的「状态」，有两个，也就是说我们需要⼀个⼆维 dp 数组。
	dp[i][j] 的定义如下：若只使⽤前 i 个物品，当背包容量为 j 时，有 dp[i][j] 种⽅法可以装满背包。
	换句话说，翻译回我们题⽬的意思就是：若只使⽤ coins 中的前 i 个硬币的⾯值，若想凑出⾦额 j，有 dp[i][j] 种凑法。
	经过以上的定义，可以得到：
	base case 为 dp[0][..] = 0， dp[..][0] = 1。因为如果不使⽤任何硬币⾯值，就⽆法凑出任何⾦额；如果凑出的⽬标⾦额为 0，
	那么“⽆为⽽治”就是唯⼀的⼀种凑法。
	我们最终想得到的答案就是 dp[N][amount]，其中 N 为 coins 数组的⼤⼩。
	第三步:根据「选择」，思考状态转移的逻辑。
	注意，我们这个问题的特殊点在于物品的数量是⽆限的，所以这⾥和之前写的背包问题⽂章有所不同。
	如果你不把这第 i 个物品装⼊背包，也就是说你不使⽤ coins[i] 这个⾯值的硬币，那么凑出⾯额 j 的⽅法数 dp[i][j] 应该等于
	dp[i-1][j]，继承之前的结果。
	如果你把这第 i 个物品装⼊了背包，也就是说你使⽤ coins[i] 这个⾯值的硬币，那么 dp[i][j] 应该等于 dp[i][j-coins[i-1]] 。
	⾸先由于 i 是从 1 开始的，所以 coins 的索引是 i-1 时表⽰第 i 个硬币的⾯值。dp[i][j-coins[i-1]] 也不难理解，
	如果你决定使⽤这个⾯值的硬币，那么就应该关注如何凑出⾦额 j - coins[i-1]。
	⽐如说，你想⽤⾯值为 2 的硬币凑出⾦额 5，那么如果你知道了凑出⾦额 3 的⽅法，再加上⼀枚⾯额为 2 的硬币，不就可以凑出 5 了嘛。
	综上就是两种选择，⽽我们想求的 dp[i][j] 是「共有多少种凑法」，所以 dp[i][j] 的值应该是以上两种选择的结果之和：
	for (int i = 1; i <= n; i++) {
		for (int j = 1; j <= amount; j++) {
			if (j - coins[i-1] >= 0)
				dp[i][j] = dp[i - 1][j] + dp[i][j-coins[i-1]];
	return dp[N][W]
	最后⼀步:把伪码翻译成代码，处理⼀些边界情况。
*/
func CompleteKnapsackProblem(amount int, coins []int) int {
	n := len(coins) // 面值不同的货币种类数
	// dp 数组初始化
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	// base case
	for i := 0; i <= amount; i++ {
		// 无硬币可选时，装法为0
		dp[0][i] = 0
	}
	for i := 0; i <= n; i++ {
		// 要凑出总额为0时，不选择便是唯一的选择
		dp[i][0] = 1
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if (j - coins[i-1]) >= 0 {
				// 该面值的硬币可供本次选择
				/*
					当能使用该面值的硬币时，总共的凑法便是不使用该硬币凑出目标金额的凑法和使用该面值硬币，
					凑出目标金额减去该硬币面值的凑法之和。前者易于理解，后者可理解为再加一枚该面值的硬币便可凑出
					目标金额，不过仍属于同一种凑法。
				*/
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			} else {
				// 该面值的硬币不能选择时，继承之前的结果
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	// dp[n][amount]：n种面额的硬币凑出金额amount总共有dp[n][amount]种凑法
	for i := 0; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			fmt.Printf("%d\t", dp[i][j])
		}
		fmt.Println()
	}
	return dp[n][amount]
}

//⽽且，我们通过观察可以发现，dp 数组的转移只和 dp[i][..] 和 dp[i- 1][..] 有关，所以可以压缩状态，进⼀步降低算法的空间复杂度：
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

/*====================================== 动态规划之背包问题 end ============================================*/

/*
	编辑距离问题就是给我们两个字符串 s1 和 s2 ，只能⽤三种操作，让我们把 s1 变成 s2 ，求最少的操作数。
	注意：解决两个字符串的动态规划问题，一般都是用两个指针 i, j 分别指向两个字符串的最后，然后一步步往前走，缩小问题规模。
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
	return MinIntAB(a, MinIntAB(b, c))
}

/*
「力扣」第 53 题（最大子序和）
给你一个整数数组 nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组:是数组中的一个连续部分。
*/
func MaxSubArray53(nums []int) int {
	// 数组长度
	length := len(nums)
	// 定义状态： dp[i] 表示以i结尾的连续子序列的最大和
	var dp = make([]int, length, length)
	// 所求最大和
	var res int
	// 初始化状态
	dp[0] = nums[0]

	for i := 1; i < length; i++ {
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
「力扣」第 70 题（爬楼梯）（23年西交915算法考题）
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
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
给你一个整数数组 coins，表示不同面额的硬币；以及一个整数 amount，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
*/
func coinChange322(coins []int, amount int) int {
	// dp[i] = x ：当⽬标⾦额为 i 时，⾄少需要 x 枚硬币
	// 数组大小为 amount+1，初始值也为 amount+1
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
			dp[i] = MinIntAB(dp[i], 1+dp[i-v])
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
给你一个整数数组 nums，找到其中最长严格递增子序列的长度。
子序列：是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
时间复杂度：O(N*N)
进阶版——时间复杂度：O(NlogN)——patience sorting——了解即可
*/
/*
	最⻓递增⼦序列（Longest Increasing Subsequence，简写 LIS）是⽐较经典的⼀个问题，⽐较容易想到的是动态规划解法，时间复杂度 O(N^2)，
	我们借这个问题来由浅⼊深讲解如何写动态规划。⽐较难想到的是利⽤⼆分查找， 时间复杂度是 O(NlogN)，我们通过⼀种简单的纸牌游戏来辅助理解这种巧妙的解法。
	注意：「⼦序列」和「⼦串」这两个名词的区别，⼦串⼀定是连续的，⽽⼦序列不⼀定是连续的。
*/
/*
卧槽，有次面试就问到了这个问题，不同的是还要求出具体的子序列，我他妈没有答出来（2022/6/28 20:18）
*/
func lengthOfLIS300(nums []int) int {
	// 序列长度
	n := len(nums)
	// step1：定义状态
	// 定义dp数组，dp[i] 表⽰以 nums[i] 这个数结尾的最⻓递增⼦序列的长度
	// 注：dp[i]：仅代表以 i 结尾的数组中最长的递增子序列的长度，又因为子序列不要求连续，故不能保证此时的长度便是以 i 结尾的数组中
	// 最长的递增子序列。
	dp := make([]int, n)
	// 初始化为1
	for e := range dp {
		dp[e] = 1
	}
	// 结果初始化为0
	res := 0

	// step2：状态转移
	// 仔细体会这个状态转移方程（2023/10/13 15:36）
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = MaxIntAB(dp[i], dp[j]+1)
			}
		}
	}

	// step3：取最优解
	for _, v := range dp {
		if res < v {
			res = v
		}
	}
	return res
}

/*
「力扣」第 416 题（分割等和子集问题）（转化为 “01背包问题” ）
时间复杂度：O(N*SUM/2)
空间复杂带：O(N*SUM/2)
*/
// 输⼊⼀个集合，返回是否能够分割成和相等的两个⼦集
func canPartition416(nums []int) bool {
	// 1.边界值处理
	// 1.1 集合不足两个元素时无法分割
	if len(nums) < 2 {
		// 元素个数小于2，则不能分割
		return false
	}
	sum := 0       // 当前集合元素总和
	max := nums[0] // 记录集合的最大值
	for _, v := range nums {
		sum += v
		if max < v {
			max = v
		}
	}
	// 1.2 和为奇数时，不可能分割成两个和相等的集合
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	// 1.3 集合最大值大于总和的二分之一时，亦不可能分割成两个和相等的集合
	if max > sum {
		return false
	}

	// 2.动态规划
	// 2.1 状态选择（变量：集合元素 + 集合总和/2）
	n := len(nums)
	// dp[i][j]：当选用前i个元素时，是否可以选则其中某些元素，使之其和为j（如何描述其子问题）
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, sum+1)
		// base case（不选取任何元素便可以凑出和为0，故dp[i][0]=true）
		dp[i][0] = true
	}
	// 2.2 状态转移（加入或不加入）
	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if (j - nums[i-1]) < 0 {
				// 当前元素值大于想要凑出的总和，故不能加入，直接继承不选当前元素的结果
				dp[i][j] = dp[i-1][j]
			} else {
				// 装入结果：若装入当前元素可以凑出想要的总和，则为true；亦或是不装入当前元素也可以凑出想要的总和，则为true
				// 装入当前元素后，只要dp[i-1][j-nums[i-1]]=true（用前i-1个元素可以凑出j-nums[i-1]）,则装入元素i之后，也可凑出j。
				dp[i][j] = dp[i-1][j-nums[i-1]] || dp[i-1][j]
			}
		}
	}
	// 2.3 取最终结果
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

/*======================================动态规划 end============================================*/
