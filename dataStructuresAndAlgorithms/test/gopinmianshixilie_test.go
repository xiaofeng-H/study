package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	var n int = 50
	counts := algorithms.CountPrimes(n)
	fmt.Printf("%d 以内总共有 %d 个素数！\n", n, counts)
}
