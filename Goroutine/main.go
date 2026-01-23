package main

import (
	"fmt"
	"sync"
)

func main() {
	//Smt()

	var money int
	var moneyCount int

	mutex := &sync.Mutex{}

	go func() {
		for {
			mutex.Lock()
			m := money
			mc := moneyCount
			mutex.Unlock()

			if m != mc {
				fmt.Printf("money: %b, money count: %b", m, mc)
				break
			}
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for range 1000 {
		go func() {
			defer wg.Done()
			mutex.Lock()
			money++
			moneyCount++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(money)
}
