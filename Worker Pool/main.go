package main

import (
	"sync"
)

func main() {

	urls := make([]string, 100000000)

	codes := make(map[int]int)
	wg := &sync.WaitGroup{}
	mx := sync.Mutex{}

	ch := make(chan string)

	go func() {
		for _, i := range urls {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(16)
	for range 16 {
		go func() {
			defer wg.Done()
			for url := range ch {
				code := httpRequest(url)

				mx.Lock()
				codes[code]++
				mx.Unlock()
			}
		}()
	}

	wg.Wait()

}

func httpRequest(url string) (code int) {
	return
}
