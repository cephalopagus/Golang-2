package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// func main() {
// 	wg := &sync.WaitGroup{}

// 	counter := 20
// 	wg.Add(counter)
// 	for i := 0; i < counter; i++ {

// 		go func() {
// 			fmt.Println(i * i)
// 			wg.Done()
// 		}()

// 	}
// 	wg.Wait()
// 	time.Sleep(time.Second)
// }

// func main() {

// 	mx := sync.Mutex{}

// 	writes := 1000
// 	storage := make(map[int]int, writes)
// 	wg := sync.WaitGroup{}
// 	wg.Add(writes)

// 	for i := 0; i < writes; i++ {
// 		go func() {
// 			mx.Lock()
// 			storage[i] = i
// 			mx.Unlock()
// 			defer wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(storage)
// }

// func main() {
// 	storage := make(map[int]int, 1000)

// 	wg := sync.WaitGroup{}
// 	reads := 1000
// 	writes := 1000
// 	mu := sync.RWMutex{}

// 	wg.Add(writes)
// 	for i := 0; i < writes; i++ {
// 		i := i
// 		go func() {
// 			defer wg.Done()

// 			mu.Lock()
// 			defer mu.Unlock()
// 			storage[i] = i
// 		}()
// 	}
// 	wg.Add(reads)
// 	for i := 0; i < reads; i++ {
// 		i := i
// 		go func() {
// 			defer wg.Done()

// 			mu.RLock()
// 			_, _ = storage[i]
// 			mu.RUnlock()
// 		}()
// 	}

// 	wg.Wait()
// 	fmt.Println(storage)
// }

func main() {
	alreadyStored := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}
	// 3, 4, 5, 0, 4, 9, 8, 6, 6, 5, 5, 4, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			mu.Lock()
			defer wg.Done()
			if _, ok := alreadyStored[doubles[i]]; !ok {

				alreadyStored[doubles[i]] = struct{}{}

				uniqueIDs <- doubles[i]
			}
			mu.Unlock()
		}()
	}

	close(uniqueIDs)
	for val := range uniqueIDs {
		fmt.Println(val)
	}
	wg.Wait()
	fmt.Printf("len of ids: %d\n", len(uniqueIDs)) // 0, 1, 2, 3, 4 ...
	fmt.Println(uniqueIDs)
}
