package main

import (
	"fmt"
	"testing"
)

type Person struct {
	age  int
	name string
}

func TestPrint(t *testing.T) {
	var age int
	var name string
	age = 26
	name = "xiaofeng"
	fmt.Printf("%v\n", age)
	fmt.Printf("%v\n", name)
	fmt.Printf("%#v\n", age)
	fmt.Printf("%#v\n", name)
	fmt.Printf("%s\n", name)

	var xiaofeng = Person{
		age:  17,
		name: "nianqi",
	}
	fmt.Printf("%+v\n", xiaofeng)
}
