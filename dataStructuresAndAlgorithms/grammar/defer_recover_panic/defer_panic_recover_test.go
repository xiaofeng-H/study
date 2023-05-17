package defer_recover_panic

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

// 函数退出前，按照先进后出的顺序，执行defer函数
func Test2(t *testing.T) {
	// defer：延迟函数执行，先进后出
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	defer fmt.Println("defer4")
	fmt.Println("11111")
}

// panic后的defer函数不会被执行（遇到panic，如果没有捕获错误，函数会立刻终止）
func Test3(t *testing.T) {
	// panic后的defer函数不会被执行
	defer fmt.Println("panic before")
	panic("发生panic")
	defer func() {
		fmt.Println("panic after")
	}()
}

// panic没有被recover时，抛出的panic到当前goroutine最上层函数时，最上层程序直接异常终止
func Test4(t *testing.T) {
	// 子函数抛出的panic没有recover时，上层函数时，程序直接异常终止
	defer func() {
		fmt.Println("c")
	}()
	F1()
	fmt.Println("继续执行")
}
func F1() {
	defer func() {
		fmt.Println("b")
	}()
	panic("a")
	fmt.Println("after F1 panic")
}

// panic有被recover时，当前goroutine最上层函数正常执行
func Test5(t *testing.T) {
	defer func() {
		fmt.Println("c")
	}()
	F2()
	fmt.Println("继续执行")
}
func F2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("b")
	}()
	panic("a")
}

// defer可以捕获到其Goroutine的子Goroutine 的panic吗?
func Test6(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("父协程捕获到子协程异常")
			fmt.Println(err)
		}
	}()
	go func() {
		//defer func() {
		//	err := recover()
		//	if err != nil {
		//		fmt.Println("子协程捕获到子协程异常")
		//		fmt.Println(err)
		//	}
		//}()
		fmt.Println("运行子协程")
		panic("子协程panic")
	}()
	fmt.Println("父协程运行")
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
}
