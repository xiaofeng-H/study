package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

/**
 *
 * 算法思想：
 * 稳定性：不稳定排序
 * 时间复杂度：O(n^2)---最坏情况：n(n-1)/2;最好情况：O(n)
 * 空间复杂度：O(1)
 * 适用场景：
 */

func TestDirectInsertSort(t *testing.T) {
	var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96,
		147, 852}
	//var b = []int{9, 8, 4, 7, 6, 2, 5, 3, 1}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.DirectInsertSort(a)
}

func TestHalfInsertSort(t *testing.T) {
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.HalfInsertSort(a)
}

func TestBubbleSort(t *testing.T) {
	var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	//var b = []int{9, 8, 4, 7, 6, 2, 5, 3, 1}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.BubbleSort(a)
}

func TestQuickSort(t *testing.T) {
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 23, 92, 87, -3, 77, -5}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.QuickSort(a, 0, len(a)-1)
}

func TestSimpleSelectSort(t *testing.T) {
	var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1,23,92,87,-3,77,-5}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.SimpleSelectSort(a)
}

func TestHeapSort(t *testing.T) {
	var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1,23,92,87,-3,77,-5}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.HeapSort(a)
}

func TestMergeSort1(t *testing.T) {
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	var a = []int{97, 65, 49, 38, 76, 13, 27, 22}
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1,23,92,87,-3,77,-5}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(a)
	algorithms.MergeSort1(a, 0, len(a)-1)
}

func TestMergeSort2(t *testing.T) {
	//var nums = []int{9, 8, 4, 7, 6, 2, 5, 3, 1, 45, 54, 98, 63, 25, 54, 885, 23687, 15, 2654, 2656, 31454, -8, -522, 96, 147, 852}
	var nums = []int{97, 65, 49, 38, 76, 13, 27, 22}
	var tmp = make([]int, len(nums))
	//var a = []int{9, 8, 4, 7, 6, 2, 5, 3, 1,23,92,87,-3,77,-5}
	fmt.Println("排序前的结果为：")
	algorithms.PrintArray(nums)
	algorithms.MergeSort2(nums, tmp, 0, len(nums)-1)
}
