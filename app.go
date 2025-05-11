package main

import (
	"fmt"
	"os"
	"strconv"
)

var filePath = "data.txt"

func writeToFile(data string){
	os.WriteFile(filePath,[]byte(data),0644)
}

func readFromFile(){
	byteData, _ := os.ReadFile(filePath)
	text := string(byteData)
	num, _ := strconv.ParseInt(text,10,64)
	fmt.Println(num)
}

func main(){
	var val string;

	fmt.Print("Enter a value :")
	fmt.Scan(&val)

	writeToFile(val)
	readFromFile()

}

