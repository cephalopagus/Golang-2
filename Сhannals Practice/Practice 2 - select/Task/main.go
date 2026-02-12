package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randTime() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
}

func predictableTime() {
	ch := make(chan struct{})
	go func() {
		randTime()
		close(ch)
	}()
	select {
	case <-ch:
		fmt.Println("прошло меньше 3 сек")
	case <-time.After(3 * time.Second):
		fmt.Println("3 секунды прошло")
		return
	}
}
func main() {
	predictableTime()
}
