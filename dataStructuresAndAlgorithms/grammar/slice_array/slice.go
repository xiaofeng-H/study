package slice_array

import "fmt"

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

func changeSlice(list []int) {
	fmt.Println("------Start the changSlice function------")
	fmt.Printf("切片变量的值为：%p,切片变量的地址为：%p\n", list, &list)
	list[len(list)-1] = len(list) * 100
	list = append(list, 3, 4, 5)
	fmt.Printf("切片变量的值为：%p,切片变量的地址为：%p\n", list, &list)
	fmt.Println("------End the changSlice function------")
}
