package main

import "fmt"

func main(){
	arr0 := make([]string,2) //make a string arr with 2 null values
	arr := make([]string,2,5) //make a string arr with 2 null values and 5 capacity
	normalArr := []string{}

	arr[0] = "hi" //works
	// normalArr[0] = "lol" //Error

	fmt.Println(arr0)
	fmt.Println(arr)
	fmt.Println(normalArr)

	//make for maps
	map1 := make(map[string]float64,4) 
	map1["car"] = 4

	fmt.Println(map1)

	for index,value := range arr{
		fmt.Println(index)
		fmt.Println(value)
	}
}