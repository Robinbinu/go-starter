package main

import (
	"fmt"
)

func main(){
	add(1,2)
	add(1.0,3.1)
	add("hello ","Hi")
}

func add[T int|float64|string](a, b T) T {
	fmt.Println(a+b)
	return a + b
}