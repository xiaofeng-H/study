package dataStructures

import "fmt"

// MoveZeroElements 将一个数组中的非零元素移动到最前后
func MoveZeroElements(a []int) {
	// 双指针
	i := 0
	j := 0
	for i < len(a) {
		if a[i] == 0 {
			i++
		} else {
			a[j] = a[i]
			i++
			j++
		}
	}
	a = a[:j]
	fmt.Println(a)
}
