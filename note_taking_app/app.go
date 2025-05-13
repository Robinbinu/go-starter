package main

import (
	"errors"
	"fmt"
	"time"
)

type Note struct{
	title string
	content string
	timeCreated time.Time
}

func main(){

	title,_:=getUserInput("Note title: ")
	content,_:=getUserInput("Note content: ")
	fmt.Println(title)	
	fmt.Print(content)						
}

func getUserInput(text string) (value string, err error){
	fmt.Print(text)
	fmt.Scanln(&value)
	if value == ""{
		return "",errors.New("invalid input")
	}
	return value,nil
}