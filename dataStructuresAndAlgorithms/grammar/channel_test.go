package grammar

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	ch := make(chan int) //这里就是创建了一个channel，这是无缓冲管道注意
	go func() { //创建子go程
		for i := 0; i < 6; i++ {
			time.Sleep(2)
			ch <- i //循环写入管道
			fmt.Printf("写入%d time:%v\n", i, time.Now().Unix())
		}
	}()

	for i := 0; i < 6; i++ { //主go程
		time.Sleep(2)
		num := <-ch //循环读出管道
		fmt.Printf("读出%d time:%v\n", num, time.Now().Unix())
	}
}

func TestChannelBuffered1(t *testing.T) {
	jobs := make(chan int, 1)
	done := make(chan bool)
	go func() {
		//		fmt.Println("GoStart")
		for i := 1; ; i++ {
			//			fmt.Println("GoforSTART", i)
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
			//			fmt.Println("GoforEND", i)
		}
	}()
	for j := 1; j <= 3; j++ {
		//		fmt.Println("OutFOR", j)
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func TestChannelBuffered2(t *testing.T) {
	// 注意：无缓冲的信道在取消息和存消息的时候都会挂起当前的goroutine，除非另一端已经准备好。
	ch := make(chan int)

	go func() {
		if element, ok := <-ch; ok {
			fmt.Println("receiver channel ", element)
		} else {
			fmt.Println("receive all elements")
		}
	}()

	// 注意：在对无缓冲区的channel发送数据的时候，必须要有一个等待接收该channel数据的协程，
	// 否则会造成死锁！
	tmp := 1
	ch <- tmp
	time.Sleep(1 * time.Second)
}

func TestChannelCommunication(t *testing.T) {
	ch := make(chan int, 3)
	go send(ch)
	time.Sleep(3 * time.Second)
	go receive(ch)
	time.Sleep(3 * time.Second)
}

func send(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("running send goroutine ", i)
	}
}

func receive(ch chan int) {
	for i := 0; i < 5; i++ {
		x := <-ch
		fmt.Println("running receive goroutine ", x)
	}
}

func TestCommunication(t *testing.T) {
	m := make(map[int]string, 5)

	go func(m map[int]string) {
		for i := 0; i < 5; i++ {
			if _, ok := m[i]; ok {
				continue
			}
			m[i] = strconv.FormatInt(int64(i), 10)
		}
		for k, v := range m {
			fmt.Printf("k = %d, v = %s\t", k, v)
		}
		fmt.Println("map[0]=", m[0])
	}(m)

	go func(m map[int]string) {
		for i := 0; i < 2; i++ {
			j := strconv.FormatInt(int64(i), 10)
			j = j + j
			m[i] = j
		}
		for k, v := range m {
			fmt.Printf("k = %d, v = %s\t", k, v)
		}
		fmt.Println("map[0]=", m[0])
	}(m)

	time.Sleep(3 * time.Second)
	for k, v := range m {
		fmt.Printf("k = %d, v = %s\t", k, v)
	}
	fmt.Println("map[0]=", m[0])
}
