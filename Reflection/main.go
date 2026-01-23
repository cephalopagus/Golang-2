package main

import (
	"fmt"
	"reflect"
)

func main() {

	Inspect(12, 22, 21)
	Inspect("Strochka")
}

func Inspect(i ...any) {
	typeR := reflect.TypeOf(i)
	valueR := reflect.ValueOf(i)

	fmt.Println(typeR, valueR, typeR.Kind())
}
