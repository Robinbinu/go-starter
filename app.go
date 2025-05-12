package main

import (
	"fmt"
	"example.com/fileops"
)



func main(){

	// panic("First panic")

	var val string;

	fmt.Print("Enter a value :")
	fmt.Scan(&val)

	fileops.WriteToFile(val)
	err:=fileops.ReadFromFile()
	if err != nil{
		panic(err)
	}

}

