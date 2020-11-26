## defer
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

- panic内置函数停止当前goroutine的正常执行，当函数F调用panic时，函数F的正常执行被立即停止，然后运行所有在F函数中的defer函数，然后F返回到调用他的函数对于调用者G，F函数的行为就像panic一样，终止G的执行并运行G中所defer函数，此过程会一直继续执行到goroutine所有的函数
- 在当前goroutine中可以使用recover捕获错误

## error 类型

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