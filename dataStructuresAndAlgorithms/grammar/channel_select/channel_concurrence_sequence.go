package main

import (
	"fmt"
	"os"
	"os/signal"
)

// 使用channel的阻塞机制，来实现goroutine顺序执行
func ConcurrenceSequence() {
	var chan1 chan int = make(chan int, 1)
	var chan2 chan int = make(chan int, 1)
	var chan3 chan int = make(chan int, 1)
	fmt.Println("ConcurrenceSequence start")
	chan1 <- 1
	go f1(chan1, chan2)
	go f2(chan2, chan3)
	go f3(chan3)

	// 阻塞主线程
	ch := make(chan os.Signal, 0)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	fmt.Println("ConcurrenceSequence end")
}

func f1(in <-chan int, out chan<- int) {
	<-in
	fmt.Println("f1 running")
	out <- 1
}

func f2(in <-chan int, out chan<- int) {
	<-in
	fmt.Println("f2 running")
	out <- 1
}
func f3(in <-chan int) {
	<-in
	fmt.Println("f3 running")
}
