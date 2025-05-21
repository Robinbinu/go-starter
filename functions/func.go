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
	fmt.Println(transformNumbers(numbers,func(i int) int {return i*4})) //[12 8 12 16]
	
	//factory function
	double := createTransformer(2) //creates a number doubler func
	fmt.Println(transformNumbers(numbers,double))
	
	fourTime := createTransformer(4)
	fmt.Println(transformNumbers(numbers,fourTime))

	//variadic function test
	sumUp("Life",2,3,45,34,1,123,343)
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

func createTransformer(factor int) func(int)int{
	return func(number int) int {
		return number*factor
}
}

 //variadic function
func sumUp(text string,numbers ...int){  //sumUp("Life",2,3,45)
	fmt.Println(text)
	for _,v := range numbers{
		fmt.Println(v)
	}
}
