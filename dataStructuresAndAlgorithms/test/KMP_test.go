package test

import (
	"fmt"
	"study/dataStructuresAndAlgorithms/algorithms"
	"testing"
)

func TestGetNext(t *testing.T) {
	var pat = "ABABABB"
	next := algorithms.GetNext(pat)
	fmt.Println(next)
}

func TestKMP(t *testing.T) {
	var str = "嘿 姑娘 你已经被我所需要 请举起你的双手 放弃无谓的抵抗 否则 我将以《中华人民共和国婚姻法》正式将你拘捕 判你余生不离我半步" +
		"晓峰吖"
	var pat = "晓峰"
	kmp := algorithms.KMP(str, pat)
	fmt.Println(kmp)
}
