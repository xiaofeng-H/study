package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
	"testing"
)

func Test1(t *testing.T) {
	var arr []int = []int{1, -3, 4, 5, -7, 10}
	res := solution(arr)
	fmt.Println(res)
}

func solution(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	len := len(arr)
	var a []int = make([]int, len)
	a[0] = arr[0]
	var res int = 0

	for i := 1; i < len; i++ {
		res += arr[i]
		if arr[i] < 0 {
			a[i] = res - arr[i]
		} else {
			a[i] = res
		}
	}

	max := 0
	for i := 0; i < len; i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	return max
}

func TestMoveZeroElements(t *testing.T) {
	var a []int = []int{0, 0, 0, 1, 0, 0, 2, 0, 3, 4, 0, 0, 5, 0,0,0,0,0,6,4,3,5,56,7,8}
	dataStructures.MoveZeroElements(a)
	fmt.Println(a)
}
