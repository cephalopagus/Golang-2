package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case val := <-ch:
		fmt.Println(val)
	default:
		fmt.Println("channal is empty")
	}
}
