package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("i =", i)

		}
		fmt.Println(runtime.NumGoroutine())
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Finish")

}
