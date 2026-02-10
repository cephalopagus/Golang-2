package main

import (
	"fmt"
	"time"
)

func main() {
	reader(double(writer()))
}

func writer() <-chan int {
	ch := make(chan int)
	go func() {
		for i := range 10 {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func double(chOld <-chan int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := range chOld {
			time.Sleep(500 * time.Millisecond)
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func reader(chOld <-chan int) {

	for i := range chOld {
		fmt.Println(i)
	}

}
