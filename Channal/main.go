package main

import "fmt"

func writer() chan int {
	ch := make(chan int)
	go func() {

		defer close(ch)

		for i := range 10 {
			ch <- i + 1
		}
		// ch <- 1
		// ch <- 2
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
}
