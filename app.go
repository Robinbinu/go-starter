package main

import (
	"fmt"
	// "example.com/fileops"
)



func main(){

	age := 32
	
	agePointer := &age //pointer referencing the address of age
	fmt.Println(agePointer," ",&age) //output: 0xc000010100   0xc000010100
	ageFromPointer := *agePointer //dereferencing the value from the pointer
	fmt.Println(ageFromPointer," ",*agePointer) //output: 32   32

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

