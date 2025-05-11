package main

import (
	"fmt"
)



func main(){

	// panic("First panic")

	var val string;

	fmt.Print("Enter a value :")
	fmt.Scan(&val)

	writeToFile(val)
	err:=readFromFile()
	if err != nil{
		panic(err)
	}

}

