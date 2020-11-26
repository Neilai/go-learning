package main

import "fmt"

func f() (result int) {
	// 函数return前执行，函数返回值为1
	defer func() {
		result++
	}()
	return 0
}

func main() {
	fmt.Println(f())
}
