package main

import (
	"fmt"
	"sync"
)

// 蓝色星际面试题：两个goroutine顺序打印数字
func main() {
	var ch1 = make(chan struct{}, 1)
	var ch2 = make(chan struct{}, 1)
	ch1 <- struct{}{}
	group := sync.WaitGroup{}
	group.Add(2)
	go grPrint1(ch1, ch2, &group)
	go grPrint2(ch1, ch2, &group)
	group.Wait()
}

func grPrint1(ch1 chan struct{}, ch2 chan struct{}, group *sync.WaitGroup) {
	defer func() {
		group.Done()
	}()
	var i int = 1
	for {
		select {
		case <-ch1:
			fmt.Printf("%d ", i)
			i += 2
			ch2 <- struct{}{}
		}
		if i > 9 {
			return
		}
	}
}

func grPrint2(ch1 chan struct{}, ch2 chan struct{}, group *sync.WaitGroup) {
	defer func() {
		group.Done()
	}()
	var i int = 2
	for {
		select {
		case <-ch2:
			fmt.Printf("%d ", i)
			i += 2
			ch1 <- struct{}{}
		}
		if i > 10 {
			return
		}
	}
}
