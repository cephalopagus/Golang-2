package main

import "fmt"

// вычисляем квадрат числа
func square(num_ch chan int) {

	value := <-num_ch       // получаем данные из main
	num_ch <- value * value // отправляем в main квадрат числа
}

// вычисляем куб числа
func cube(num_ch chan int) {

	value := <-num_ch               // получаем данные из main
	num_ch <- value * value * value // отправляем в main куб числа
}

func main() {

	// создаем два канала
	sqr_ch := make(chan int)
	cube_ch := make(chan int)

	// запускаем две горутины для взаимодействия через каналы
	go square(sqr_ch)
	go cube(cube_ch)

	// отправляем данные в каналы
	sqr_ch <- 3
	cube_ch <- 5

	// обработка данных
	select {

	case sqr_val := <-sqr_ch:
		fmt.Println("Square:", sqr_val) // 9

	case cube_val := <-cube_ch:
		fmt.Println("Cube:", cube_val) // 125
	}
}
