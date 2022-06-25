package grammar

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	arr1 := make([]int, 0, 10)
	arr2 := make([]int, 5, 10)
	fmt.Printf("The length of arr1 = %d\t, The capacity of arr1 = %d\t\n", len(arr1), cap(arr1))
	fmt.Printf("The length of arr2 = %d\t, The capacity of arr2 = %d\t\n", len(arr2), cap(arr2))
	arr1 = append(arr1, 1)
	arr2 = append(arr2, 2)
	fmt.Printf("The length of arr1 = %d\t, The capacity of arr1 = %d\t\n", len(arr1), cap(arr1))
	fmt.Printf("The length of arr2 = %d\t, The capacity of arr2 = %d\t\n", len(arr2), cap(arr2))
}

func Test2(t *testing.T) {
	arr := make([]int, 0)
	fmt.Printf("The length of arr = %d\t, The capacity of arr = %d\t\n", len(arr), cap(arr))
	arr = append(arr, 1)
	fmt.Printf("The value of arr[0] is %d BEFORE change\n", arr[0])
	change1(arr)
	fmt.Printf("The value of arr[0] is %d AFTER change\n", arr[0])
	for e := range arr {
		fmt.Println(arr[e])
	}
	change2(&arr)
	for e := range arr {
		fmt.Println(arr[e])
	}
}

func change1(arr []int) {
	arr[0] = 0
	arr = append(arr, 2)
	arr = append(arr, 3)
	for e := range arr {
		fmt.Println(arr[e])
	}
}

func change2(arr *[]int) {
	*arr = append(*arr, 2)
	*arr = append(*arr, 3)
}
