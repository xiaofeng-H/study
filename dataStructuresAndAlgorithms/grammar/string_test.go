package grammar

import (
	"fmt"
	"testing"
)

func TestStringToArray(t *testing.T) {
	var str string
	str = "oxAAtyyu"
	data := []byte(str)
	fmt.Printf("%v\n", data[0])
	fmt.Println(data[0]=='o')
	fmt.Println(string(str[4]))
	fmt.Println(data[1])
}
