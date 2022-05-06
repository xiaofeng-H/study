package algorithms

import "fmt"

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
		if isPrime[i] {
			for j := i * i; j < n; j += i {
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
算法名称：
算法思想：
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
	max := -1
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
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
