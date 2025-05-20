package main

import (
	"fmt"
)
type transformedFn func(int)int

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println(transformNumbers(numbers,getTransformFn(numbers))) //[2 4 6 8]
	numbers[0] = 3
	fmt.Println(transformNumbers(numbers,getTransformFn(numbers))) //[9 6 9 12]
}

func transformNumbers(nums []int, transform transformedFn) []int {
	transformedNums := []int{}
	for _, val := range nums {
		transformedNums = append(transformedNums, transform(val))
	}
	return transformedNums
}

func getTransformFn(nums []int) transformedFn{
	if nums[0]==1{
		return double
	}else{
		return triple
	}
}

func double(num int) int {
	return num * 2
}


func triple(num int) int {
	return num * 3
}
