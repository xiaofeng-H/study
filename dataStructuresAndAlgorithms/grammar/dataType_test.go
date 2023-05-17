package main

import (
	"fmt"
	"testing"
)

func TestUint(t *testing.T) {
	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)
}

func TestRune(t *testing.T) {
	var str string = "string字符串"
	var strArr1 []rune = []rune(str)
	var strArr2 []byte = []byte(str)
	fmt.Println(len(str))
	fmt.Println(len(strArr1))
	fmt.Println(len(strArr2))
	fmt.Println(string(str[8]))
	fmt.Println(string(strArr1[8]))
	fmt.Println(string(strArr2[8]))
}
