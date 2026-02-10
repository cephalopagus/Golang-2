package main

import (
	"fmt"
	"sync"
	"time"
)

func writer() chan int { //(<-chan int) - если хочу вернуть канал только на чтение, но в самом методе он доступен и для записи
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 1
		}
		// ch <- 1
		// ch <- 2
	}()

	go func() {
		defer wg.Done()
		for i := range 5 {
			ch <- i + 100
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {

	ch := writer()

	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	time.Sleep(1 * time.Second)
}
