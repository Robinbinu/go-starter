package main

import(
    "fmt"
	"os"
	"strconv"
)

var filePath = "data.txt"

func writeToFile(data string){
	os.WriteFile(filePath,[]byte(data),0644)
}

func readFromFile() (err error){
	byteData, err := os.ReadFile(filePath) //error handling
	if err != nil{
		return
	}
	text := string(byteData)
	num, _ := strconv.ParseInt(text,10,64)
	fmt.Println(num)
	return
}