package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	listPerson := []Person{
		{name: "Dastan", age: 23},
		{name: "Isma", age: 10},
		{name: "Adis", age: 27},
	}

	firstPerson := Person{
		name: "Dastan",
		age:  23,
	}

	listInt := []int{1, 2, 3, 4}
	listFloat := []float32{1.2, 34.5, 13.1, 6.0}

	fmt.Println(SumSmth(listInt))
	fmt.Println(SumSmth(listFloat))

	fmt.Println(ConPers(listPerson, firstPerson))
	fmt.Println(ConPers(listInt, 6))

}

func ConPers[T comparable](list []T, value T) bool {
	for _, i := range list {
		if value == i {
			return true
		}
	}
	return false
}

func SumSmth[T int | float32](numbers []T) T {
	var sum T

	for _, i := range numbers {
		sum += i
	}
	return sum
}
