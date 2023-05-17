package main

import (
	"fmt"
	"testing"
	"time"
)

var complete chan bool = make(chan bool)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	complete <- true // 执行完毕了，发个消息
}

func TestGoroutine1(t *testing.T) {
	go loop()
	<-complete // 直到线程跑完, 取到消息. main在此阻塞住
}

func TestGoroutine2(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	time.Sleep(3 * time.Second)
}
