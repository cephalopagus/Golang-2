package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
		}
	}()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
