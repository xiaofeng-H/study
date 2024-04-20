package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToArray(t *testing.T) {
	var str string
	str = "oxAAtyyu"
	//data := []byte(str)
	//fmt.Printf("%v\n", data[0])
	//fmt.Println(data[0] == 'o')
	//fmt.Println(string(str[4]))
	//fmt.Println(data[1])
	for k, v := range str {
		fmt.Println(k, ":", v)
		fmt.Println(k, ":", string(v))
	}
	fmt.Println(str[1:3])
	fmt.Printf("%T\n", str[0])
}

func TestStringToNum(t *testing.T) {
	str := "12345"
	fmt.Println(str[0], str[1])
	fmt.Println(str[0], str[0:])
	arr := []byte(str)
	fmt.Printf("%d--%d\n", arr[1], arr[1]+1)
	fmt.Printf("%s--%s\n", string(arr[1]), string(arr[1]+1))
}

// interface强转为其他类型
func TestInterfaceToString(t *testing.T) {
	var a, b interface{}
	a = 'a'
	b = 'a'
	fmt.Printf("a = %v; a.Type = %T\n", string(a.(rune)), a)
	fmt.Printf("b = %v; b.Type = %T\n", b, b)
	c := int(97)
	fmt.Printf("c = %v; c.Type = %v\n", c, reflect.TypeOf(c).Name())
}

func TestMaxProfit(t *testing.T) {
	var prices = []int{3, 3, 5, 0, 0, 3, 1, 4, 23, 4, 5435, 4544, 4, 5, 6, 6, 7, 45, 2, 45, 2444, 1, 67, 7, 54, 234, 1, 432, 6, 11, 3445}
	maxProfit(prices)
}

func maxProfit(prices []int) int {
	/* 题解：
	   怎么感觉可以用我解决Ⅱ用的代码呀！
	   eee，其实是不一样的。
	*/

	var days = len(prices)
	if days <= 1 {
		return 0
	}
	// dp[i]:第i天可获得的最大利润
	var dp = make([]int, days)
	dp[0] = 0
	// 已某天为一个截至日期，则在该截止日期可获得的最大利润
	var pro = 0

	// 动规（哈哈哈 都不知道算不算动规了，感觉都快成暴力了）
	var v = 0 // 当日股价
	for i := 1; i < days; i++ {
		// 局部变量初始化
		v = prices[i] // 当日股价
		pro = 0       // 利润初始值
		// 计算已i天为截至日期，则可获得的最高利润（遍历直到该天已被操作或到达第一天）
		for j := i - 1; j >= 0; j-- {
			if dp[j] > 0 {
				// j天为一个此前的一个截至日期，此时需要检测需不需要合并（求合并与否的利润最大值）
				if pro < v-prices[j] {
					// 到上一个截止日期为止（不包含该天），若可得利润小于当日股价与上一个截至日期股价的插值（即利润），
					// 则需要合并这两个小利润周期。因为dp[j]已经记录了第j天的最大利润，故不需要向前遍历。
					// 执行合并操作
					pro = v - prices[j] + dp[j] // 计算合并后的利润
					dp[j] = 0                   // 消除该利润周期
				}
				break
			}
			// 计算目前利润周期可获得的最大利润
			pro = max(pro, v-prices[j])
		}
		dp[i] = pro
	}

	for k, v := range dp {
		fmt.Printf("k:%d v:%d \n", k, v)
	}
	// 求利润和（利润最大的两笔）
	if days < 2 {
		return dp[0]
	}
	tmp := 0
	index := 0
	for i := 0; i < 2; i++ {
		tmp = dp[i]
		index = i
		for j := i; j < days; j++ {
			if dp[j] > tmp {
				tmp = dp[j]
				index = j
			}
		}
		// 交换
		dp[index] = dp[i]
		dp[i] = tmp
	}

	return dp[0] + dp[1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
