package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() {
	// 定义变量
	var (
		name    int
		age     int
		married int
	)
	ag1, err := fmt.Scan(&name, &age, &married)                         // 返回类型根据使用场景可以不写，但是根据代码规范，err都要处理，加上在这里为了更加严谨，确实需要err来处理
	fmt.Println(err)                                                    // 这里如果出错，则显示出错的原因
	fmt.Println(ag1)                                                    // 这个返回参数取到的是成功了几个
	fmt.Printf("扫描结果：name:%d age:%d married:%d \n", name, age, married) // Scan读取以空白为分割的值保存给参数中，换行符视为空白符，如果返回的数据比提供的参数少，会返回错误原因
}

func ScanDemo() {
	var arr []int
	for {
		var num int
		n, err := fmt.Scanln(&num)
		fmt.Println(n, err)
		fmt.Println(num)
		if n == 1 && err == nil {
			arr = append(arr, num)
		}
		if n == 0 {
			fmt.Println(err)
			break
		}
	}
	fmt.Println("接收到的数据：", arr)
}

func Scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())

}
