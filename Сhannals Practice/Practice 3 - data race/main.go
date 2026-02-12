// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup

// var mu sync.Mutex

// var value = 22

// func update_value(id int) {

// 	fmt.Println("Goroutine", id, "starts")
// 	if value == 22 {

// 		time.Sleep(1 * time.Second)
// 		fmt.Println("Goroutine", id, "changes the value")
// 		value += 1

// 	}

// 	fmt.Println(value)
// 	wg.Done() // сигнализируем, что горутина завершила выполнение
// }

// func main() {

// 	wg.Add(2) // ждем 2 горутины

// 	go update_value(1)
// 	go update_value(2)

// 	wg.Wait() // ожидаем завершения горутин
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0 //  общий ресурс
var wg sync.WaitGroup

var mu sync.Mutex

func work(number int) {
	mu.Lock()
	defer mu.Unlock()
	counter = 0 // сбрасываем общий ресурс

	for k := 1; k <= 5; k++ {
		counter += 1                       // изменяем общий ресурс
		time.Sleep(100 * time.Millisecond) // задержка для наглядности
		fmt.Println("Goroutine", number, "-", counter)
	}
	wg.Done() // сигнализируем, что горутина завершила работу
}

func main() {

	goroutines_count := 4 // количество запускаемых горутин

	wg.Add(goroutines_count)

	// запускаем горутины
	for i := 1; i <= goroutines_count; i++ {
		go work(i)
	}
	// ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("The End")
}
