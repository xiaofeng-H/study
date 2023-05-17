## 一、应用场景

1、协程间通信，即协程间数据传递

2、并发场景下利用channel的阻塞机制，作为同步机制（类似队列）

3、利用channel关闭时发送广播的特性，作为协程退出通知

## 二、通信共享内存

1、channel的方向：读、写、读和写

2、channel协程间通信信道

3、channel的阻塞机制

```go
// 利用channel的阻塞机制，等待退出信号
ch := make(chan os.Signal, 0)
signal.Notify(ch, os.Interrupt, os.Kill)
<-ch
```

4、channel并发场景下的同步机制

5、channel通知协程退出

6、channel多路复用

## 三、注意

1、channel死锁问题：必须存在读写双方，channel不能只有一个协程操作！