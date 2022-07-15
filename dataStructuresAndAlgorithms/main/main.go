package main

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
)

func main() {
	//algorithms.Random()
	//algorithms.Conversion()

	// Scan测试
	//grammar.Scan()
	//grammar.ScanDemo()
	//grammar.Scanner()

	// 反射测试
	//grammar.ReflectStructTest()

	fmt.Println("请输入楼梯层数")
	var tmp int
	fmt.Scan(&tmp)
	stairs70 := algorithms.ClimbStairs70(tmp)
	fmt.Printf("总共有%d种爬法！", stairs70)
}
