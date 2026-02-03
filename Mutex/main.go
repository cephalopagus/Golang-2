package main

import (
	"fmt"
	"sync"
)

const numGor = 1000

func main() {
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(numGor)
	value := 0
	for i := 0; i < numGor; i++ {
		go func() {
			mx.Lock()
			value++
			mx.Unlock()
			defer wg.Done()
		}()

	}
	wg.Wait()

	fmt.Println(value)

}
