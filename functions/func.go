package main

import (
	"fmt"
)
type transformedFn func(int)int

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(transformNumbers(numbers, double))
	fmt.Println(transformNumbers(numbers, triple))
}

func transformNumbers(nums []int, transform transformedFn) []int {
	transformedNums := []int{}
	for _, val := range nums {
		transformedNums = append(transformedNums, transform(val))
	}
	return transformedNums
}

func double(num int) int {
	return num * 2
}


func triple(num int) int {
	return num * 3
}
