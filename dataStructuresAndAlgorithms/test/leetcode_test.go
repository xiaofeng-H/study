package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestString(t *testing.T) {
	s := "ABCDE"
	for _, e := range s {
		fmt.Println(e == 'A')
	}
	fmt.Println(s + s)
}

func TestLeetCode119(t *testing.T) {
	row := algorithms.GetRow(3)
	fmt.Println(row)
}
