package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch1 <- 12
	ch1 <- 121
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := merge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

func merge[T any](chans ...chan T) chan T {
	ch := make(chan T)
	wg := &sync.WaitGroup{}
	wg.Add(len(chans))
	for _, v := range chans {
		go func() {
			for val := range v {
				ch <- val
			}
			defer wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(ch)

	}()

	return ch
}
