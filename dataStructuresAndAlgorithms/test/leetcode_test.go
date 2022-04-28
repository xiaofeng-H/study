package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestLeetCodeGetRow(t *testing.T) {
	row := algorithms.GetRow(3)
	fmt.Println(row)
}

func TestLeetCodeMinWindow(T *testing.T) {
	s := "ADOBECODEBANC"
	t := "ABC"
	str := algorithms.MinWindow(s, t)
	fmt.Println(str)
}
