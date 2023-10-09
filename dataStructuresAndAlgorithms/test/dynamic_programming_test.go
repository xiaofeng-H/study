package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

// 完全背包问题
func TestCompleteKnapsack(t *testing.T) {
	var coins = []int{1, 2, 5}
	var amount = 10
	res := algorithms.CompleteKnapsackProblem(amount, coins)
	fmt.Println(res)
}
