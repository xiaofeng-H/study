package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestGo(t *testing.T) {
	var char1 = '*'
	data := algorithms.Opt(3, 4, char1)
	fmt.Println(data)
}
