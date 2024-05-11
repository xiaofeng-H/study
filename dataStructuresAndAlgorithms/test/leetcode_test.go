package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestLeetCodeGetRow(t *testing.T) {
	row := algorithms.GetRow119(3)
	fmt.Println(row)
}

func TestLeetCodeMinWindow(T *testing.T) {
	s := "ADOBECODEBANC"
	t := "ABC"
	str := algorithms.MinWindow76(s, t)
	fmt.Println(str)
}

func TestLeetCode752(t *testing.T) {
	var deadends = []string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}
	var target = "8888"
	step := algorithms.OpenLock752(deadends, target)
	fmt.Println(step)
}

func TestLeetCode46(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	permute46 := algorithms.Permute46(arr)
	fmt.Println(permute46)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
