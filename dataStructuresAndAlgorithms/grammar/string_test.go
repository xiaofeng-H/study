package grammar

import (
	"fmt"
	"testing"
)

func TestStringToArray(t *testing.T) {
	var str string
	str = "oxAAtyyu"
	//data := []byte(str)
	//fmt.Printf("%v\n", data[0])
	//fmt.Println(data[0] == 'o')
	//fmt.Println(string(str[4]))
	//fmt.Println(data[1])
	for k, v := range str {
		fmt.Println(k, ":", v)
		fmt.Println(k, ":", string(v))
	}
	fmt.Println(str[1:3])
	fmt.Printf("%T\n", str[0])
}

func TestStringToNum(t *testing.T) {
	str := "12345"
	fmt.Println(str[0], str[1])
	arr := []byte(str)
	fmt.Printf("%d--%d\n", arr[1], arr[1]+1)
	fmt.Printf("%s--%s\n", string(arr[1]), string(arr[1]+1))
}
