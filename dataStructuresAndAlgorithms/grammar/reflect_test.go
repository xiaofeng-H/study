package grammar

import (
	"fmt"
	"reflect"
	"testing"
)

type Cat struct {
	name string
}

func (ca *Cat) eat() {
	fmt.Printf("I am cat named %s and like eating mouses\n", ca.name)
}

func TestReflect1(t *testing.T) {
	var mimi Cat = Cat{name: "mimi"}
	myCat(mimi)
}

func myCat(myCat Cat) {
	cat := reflect.ValueOf(myCat)
	field := cat.Field(0)
	fmt.Println("I am cat ", field)
}

type Student struct {
	Id   int
	Name string
}

func (s Student) Hello() {
	fmt.Println("我是一个学生")
}

func ReflectStructTest() {
	s := Student{Id: 1, Name: "咖啡色的羊驼"}

	// 获取目标对象的类型
	t := reflect.TypeOf(s)
	fmt.Printf("reflect.TypeOf()返回值的类型为：%T\n", t)
	// .Name()可以获取去这个类型的名称
	fmt.Println("这个类型的名称是:", t.Name())

	// 获取目标对象的值
	v := reflect.ValueOf(s)
	// .NumField()来获取该类型包含字段的总数
	for i := 0; i < t.NumField(); i++ {
		// 从0开始获取Student所包含的key
		key := t.Field(i)

		// 通过interface方法来获取key所对应的值
		value := v.Field(i).Interface()

		fmt.Printf("第%d个字段是：%s:%v = %v \n", i+1, key.Name, key.Type, value)
	}

	// 通过.NumMethod()来获取Student里头的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("第%d个方法是：%s:%v\n", i+1, m.Name, m.Type)
	}
}

func TestReflect2(t *testing.T) {
	var r = rune(32)
	fmt.Println(reflect.TypeOf(r).Name()=="int8")
}
