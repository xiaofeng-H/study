package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/dataStructures"
	"testing"
)

// 并查集测试
func TestDisjointSet(t *testing.T) {
	var edges = [][2]int{
		{0, 1}, {1, 2}, {1, 3}, {2, 5}, {3, 4},
	}
	var nodeCounts = 6 // 顶点个数
	// 初始化并查集
	var un *dataStructures.UnionFindSet
	un = dataStructures.NewUF(nodeCounts)

	// 合并操作
	for i := 0; i < len(edges); i++ {
		// 判断并查集中是否存在环，有环则返回，无环则合并
		if un.Connected(edges[i][0], edges[i][1]) {
			fmt.Println("Cycle is found!")
			return
		}
		un.Union(edges[i][0], edges[i][1])
	}
	fmt.Println("No cycle found!")

	for k, v := range un.GetUnionFindSet() {
		fmt.Printf("node=%d parents=%d\t", k, v)
		// 换行符
		if k == 3 {
			fmt.Println()
		}
	}
}
