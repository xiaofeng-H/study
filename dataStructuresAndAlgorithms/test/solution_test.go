package test

import (
	"fmt"
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
