package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestDisjointSet(t *testing.T) {
	var edges = [][2]int{
		{0, 1}, {1, 2}, {1, 3}, {2, 5}, {3, 4},
	}
	var n = len(edges)	// 顶点个数
	var parents []int = make([]int, 10)
	algorithms.InitialParents(parents)
	for i := 0; i < n; i++ {
		isOk := algorithms.Union(edges[i][0], edges[i][1], parents)
		if isOk == 0 {
			fmt.Println("Cycle is found!")
			return
		}
	}
	fmt.Println("No cycle found!")
	for k, v := range parents {
		fmt.Printf("index=%d value=%d, ", k, v)
	}
}
