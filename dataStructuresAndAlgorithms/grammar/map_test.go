package grammar

import (
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
	var m map[int]int
	i, ok := m[1]
	fmt.Println(i, ok)
	m[2]++
	fmt.Println(m[2])
	// 错误操作：map未初始化
	//m[2]=10
	//fmt.Println(m[2])

	m = make(map[int]int)

}
