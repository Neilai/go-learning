# Basic
## 数组和slice
[demo](https://github.com/Neilai/go-learning/blob/master/basic/function/func-demo.go)

要点：
- 数组是传值的
- 切片是引用，是数组的一个view
- 切片类似C++的vector，调用cap len可以得到它的容量和现在的长度
- slice append超过cap会造成内存重分配

## 结构体方法
[demo](https://github.com/Neilai/go-learning/blob/master/basic/method/method-demo.go)

主要是有传值和传指针两种，传指针会修改原对象。
## function
[demo](https://github.com/Neilai/go-learning/blob/master/basic/function/func-demo.go)
- 不定参数语法...
- 函数是一等公民
- 函数返回值可以写在定义处
- 多返回值，第二个通常是error/ok
## map
[demo](https://github.com/Neilai/go-learning/blob/master/basic/maps/map-demo.go)
创建，删除，取值，遍历语法见demo
## make 和 new的区别
- new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值

- make 只能用于 slice，map，channel 三种类型，make(T, args) 返回的是初始化之后的 T 类型的值，这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用

# Go 反射
[demo](https://github.com/Neilai/go-learning/blob/master/reflect/reflect-demo.go)

反射是指在运行时操作任意类型对象的能力。我们可以通过reflect.ValueOf和reflect.typeOf获取对象的值和类型。
运行时遍历对象字段和方法见demo

# Go module

## go module 使用
[demo](https://github.com/Neilai/go-learning/tree/master/modules)

- golang 升级到 1.11后可以使用 新的模块管理方式，之前是使用类似Vendor的机制，所有的包往Gopath里面找当modules 功能启用时，依赖包的存放位置变更为$GOPATH/pkg。
- goland若是识别不到go modules的包，需要在设置里面开启go module integration。
- go.mod 类似 package.json  go.sum 类似package-lock.json 
- 常用命令 go mod init(类似npm init) ,go mod tidy(类似npm install)
- 使用go mod时 , 导入本地包不要使用相对路径，要根据顶层Mod的package name往下写（和目录名无关）

## init函数及其调用顺序

init是导入的时候执行的，先调用依赖的依赖init，同层依赖init按字典序调用。


#  error Handling

## defer

[demo](https://github.com/Neilai/go-learning/blob/master/errorHandling/defer/defer-demo.go)

defer函数返回时执行,通常用于资源释放、打印日志、异常捕获。
```go
func main() {
    f, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer f.Close()
}
```

## panic  recover

[demo](https://github.com/Neilai/go-learning/blob/master/errorHandling/panic/panic-demo.go)

- panic内置函数停止当前goroutine的正常执行，当函数F调用panic时，函数F的正常执行被立即停止，然后运行所有在F函数中的defer函数，然后F返回到调用他的函数对于调用者G，F函数的行为就像panic一样，终止G的执行并运行G中所defer函数，此过程会一直继续执行到goroutine所有的函数
- 在当前goroutine中可以使用recover捕获错误

## error 类型

## panic  recover

[demo](https://github.com/Neilai/go-learning/blob/master/errorHandling/error/error-demo.go)

error类型其实是一个接口
```go
type error interface {
    Error() string
}
```
go中，函数通常在最后的返回值中返回错误信息，使用errors.New 可返回一个错误信息,erros.New实际上就是生成了一个包含错误字符串的errorStr

## golang 实现try catch
见demo,自定义Try 函数，第一个参数是要执行的函数，第二个函数放在defer中就实现了类似Java中的try catch机制

## Go错误处理的争议

Go语言推荐使用它自己的一套错误处理机制。其错误处理机制首先自定义自己的错误结构体(struct)，然后基于这个结构体实现error接口的函数 Error() 并回传一个错误信息字符串。

# interface

## Go的接口
[demo](https://github.com/Neilai/go-learning/blob/master/interface/interface-demo.go)

- go的接口是鸭子类型，弱约束,兼具python和Java的优点
- 比较奇怪的一点是Go interface若是pointer调用，只能传指针，若是value 调用，传指针和值都可以
- interface可通过 x.type判断实际调用类型,见demo

# goroutine

goroutine即协程，go内部有一个协程的逻辑调度器，它的调度是协作式的，非抢占，遇到IO(网络操作、文件读写、print)会让出事件片
## waitgroup 使用 
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/waitgroup/waitgroup-demo.go)

waitgroup 类似信号量,add操作相当于设置初始值，done()相当于 signal操作，它主要用来main函数同步其他协程，防止main退出导致所有协程终止。
## GMP 调度原理
[demo(调整GOMAXPROCS)](https://github.com/Neilai/go-learning/blob/master/goroutine/waitgroup/waitgroup-demo.go)

![](http://www.topgoer.com/static/7.1/gmp/12.jpg)

M(线程)的数量是内部设定的，Go程序会尽量多设置。

P(逻辑调度器)的数量通过GOMAXPROCS设置，通过waitgroup demo也可以看到，当runtime.GOMAXPROCS(2)设置为2的时候，两个协程是并行运行输出的。

## 竞态
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/race/race-demo.go)

多个协程操作同一数据，可能会发生竞态，在race-demo中，我们通过runtime.Gosched()主动让出时间片重现了这一过程。
## mutex/atomic
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/mutex/mutex-demo.go)

go提供了互斥锁和原子操作，用于对于竞态情况处理。
## 条件变量
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/cond/cond-demo.go)

类似其他编程语言，Go也提供了条件变量，条件变量是对mutex的补充，在有条件判断的竞争程序，可以简化程序编写，减少不必要的acquire waita操作。
cond-demo是一个生产者消费者的例子。
## unbufferdChannel
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/unbufferdChannel/unbufferdChannel.go)

Channel 类似进程通信的管道，其中unbufferdChannel没有缓冲，意味着管道接受端没有接收到数据会阻塞。
unbufferdChannel-demo是一个channel模仿打球(一来一回)的例子
## bufferdChannel
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/bufferdChannel/bufferdChannel.go)

bufferdChannel就是带有缓冲的管道，若是缓冲未满可以发送数据，buffered-demo是一个使用channel的生产者消费者例子。
## context
[demo](https://github.com/Neilai/go-learning/blob/master/goroutine/context/context-demo.go)

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

# 单元测试

## 单元测试
[demo](https://github.com/Neilai/go-learning/blob/master/unitTest/demo_1_test.go)

- 要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时需要让文件必须以_test结尾。默认的情况下，go test命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数。
- 测试用例文件使用go test指令来执行，没有也不需要 main() 作为函数入口。所有在以_test结尾的源码内以Test开头的函数会自动被执行

## 基准测试
[demo](https://github.com/Neilai/go-learning/blob/master/unitTest/demo_2_test.go)

基准测试的函数必须以Benchmark开头，必须是可导出的
b.ResetTimer是重置计时器，这样可以避免for循环之前的初始化代码的干扰
b.N是基准测试框架提供的，表示循环的次数，因为需要反复调用测试的代码，才可以评估性能
## convey
[demo](https://github.com/Neilai/go-learning/blob/master/modules/convey_test.go)

我们项目中大量使用了convey单元测试框架，demo见modules/convey-demo.go

