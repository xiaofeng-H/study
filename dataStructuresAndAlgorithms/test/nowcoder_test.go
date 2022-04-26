package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestNQueens(t *testing.T)  {
	queens := algorithms.SolveNQueens(8)
	fmt.Printf("总共%d种解！\n",len(queens))
	for i := range queens {
		for _, e := range queens[i] {
			fmt.Println(e)
		}
		fmt.Println("===============Next===============")
	}
}
