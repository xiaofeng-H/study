package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestBaseTrans(t *testing.T) {
	N := 16
	B := 2
	res := algorithms.BaseTrans(N, B)
	fmt.Printf("将【%d】转化为【%d】进制的结果是：%d\n", N, B, res)

}
