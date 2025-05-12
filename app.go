package main

import (
	"fmt"
	// "example.com/fileops"
)

func incrementAgeBy10(age *int){ //func requires a pointer param 
	*age = *age + 10 //pointer dereferenced and mutated
}


func main(){

	age := 32
	fmt.Println(age) //output: 32
	incrementAgeBy10(&age)
	fmt.Println(age) //output: 42


	// panic("First panic")

	// var val string;

	// fmt.Print("Enter a value :")
	// fmt.Scan(&val)

	// fileops.WriteToFile(val)
	// err:=fileops.ReadFromFile()
	// if err != nil{
	// 	panic(err)
	// }

}

