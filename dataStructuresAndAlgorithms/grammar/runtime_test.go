package main

import (
	"fmt"
	"runtime"
	"testing"
)

/*======================================runtime start============================================*/
/*rt.1 Gosched
runtime.Gosched() 用于让出CPU时间片，让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
这就像跑接力赛，A跑了一会碰到代码runtime.Gosched() 就把接力棒交给B了，A歇着了，B继续跑。
示例代码：*/
func TestRunTime1(t *testing.T) {
	//创建一个goroutine
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	for i := 0; i < 2; i++ {
		runtime.Gosched() //import "runtime"
		/*
		   屏蔽runtime.Gosched()运行结果如下：
		       hello
		       hello

		   没有runtime.Gosched()运行结果如下：
		       world
		       world
		       hello
		       hello
		*/
		fmt.Println("hello")
	}
}

/*
rt.2 Goexit
调用 runtime.Goexit() 将立即终止当前 goroutine 执⾏，调度器确保所有已注册 defer 延迟调用被执行。
示例代码：
*/
func TestRunTime2(t *testing.T) {
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() // 终止当前 goroutine, import "runtime"
			fmt.Println("B") // 不会执行
		}()

		fmt.Println("A") // 不会执行
	}() //别忘了()

	//死循环，目的不让主goroutine结束
	for {
		var tmp int = 0
		tmp++
		//fmt.Println(tmp)
	}
}

/*
rt.3 GOMAXPROCS
调用 runtime.GOMAXPROCS() 用来设置可以并行计算的CPU核数的最大值，并返回之前的值。
示例代码：
*/
func TestRunTime3(t *testing.T) {
	n := runtime.GOMAXPROCS(1) //打印结果：111111111111111111110000000000000000000011111...
	//n := runtime.GOMAXPROCS(2)     //打印结果：010101010101010101011001100101011010010100110...
	fmt.Printf("n = %d\n", n)

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}

/*
在第一次执行(runtime.GOMAXPROCS(1))时，最多同时只能有一个goroutine被执行。所以
会打印很多1。过了一段时间后，GO调度器会将其置为休眠，并唤醒另一个goroutine，这时候就开始打印很多0了，在打印的时候，goroutine是被调度到操作系统线程上的。
在第二次执行(runtime.GOMAXPROCS(2))时，我们使用了两个CPU，所以两个goroutine可以一起被执行，以同样的频率交替打印0和1。
*/
/*======================================runtime end============================================*/
