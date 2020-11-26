// This sample program demonstrates how to create race
// conditions in our programs. We don't want to do this.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	go incCounter()
	go incCounter()
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incCounter() {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}
