package main

import (
	"fmt"
	"sync"
	"time"
)

func Smt() {
	counter := 20

	wg := &sync.WaitGroup{}

	for i := 0; i < counter; i++ {
		wg.Add(1)
		i := i
		go func() {
			fmt.Println(i * i)
			wg.Done()
		}()

	}
	wg.Wait()

	time.Sleep(time.Second)

}
