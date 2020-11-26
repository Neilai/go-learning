## goroutine
goroutine即协程，go内部有一个协程的逻辑调度器，它的调度是协作式的，非抢占，遇到IO(网络操作、文件读写、print)会让出事件片
## waitgroup 使用 
waitgroup 类似信号量,add操作相当于设置初始值，done()相当于 signal操作，它主要用来main函数同步其他协程，防止main退出导致所有协程终止。
## GMP 调度原理
M(线程)的数量是内部设定的，Go程序会尽量多设置。
P(逻辑调度器)的数量通过GOMAXPROCS设置，通过waitgroup demo也可以看到，当runtime.GOMAXPROCS(2)设置为2的时候，两个协程是并行运行输出的。
## 竞态
多个协程操作同一数据，可能会发生竞态，在race-demo中，我们通过runtime.Gosched()主动让出时间片重现了这一过程。
## mutex/atomic
go提供了互斥锁和原子操作，用于对于竞态情况处理。
## 条件变量
类似其他编程语言，Go也提供了条件变量，条件变量是对mutex的补充，在有条件判断的竞争程序，可以简化程序编写，减少不必要的acquire waita操作。
cond-demo是一个生产者消费者的例子。
## unbufferdChannel
Channel 类似进程通信的管道，其中unbufferdChannel没有缓冲，意味着管道接受端没有接收到数据会阻塞。
unbufferdChannel-demo是一个channel模仿打球(一来一回)的例子
## bufferdChannel
bufferdChannel就是带有缓冲的管道，若是缓冲未满可以发送数据，buffered-demo是一个使用channel的生产者消费者例子。
## context
上下文，用于更方便的控制Goroutine,例如对协程的取消、传值操作。

context background和todo的区别:
```go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```

它们两个本质上都是emptyCtx结构体类型，是一个不可取消，没有设置截止时间，没有携带任何值的Context。
目前没太理解它们的区别，好像只有语义的区别？？不知道传什么的时候就传todo
