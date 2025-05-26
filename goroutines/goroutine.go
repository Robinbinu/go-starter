package main

import (
	"fmt"
	"time"
)

func main(){
	dataChan := make(chan bool)
	go greet(dataChan)
	go greet(dataChan)
	go slowGreet(dataChan)
	go greet(dataChan)

	for i:=0;i<4;i++{
		<- dataChan
	}

}

func greet(dataChan chan bool){
	fmt.Println("Hi")
	dataChan <- true
}

func slowGreet(dataChan chan bool){
	time.Sleep(3*time.Second)
	fmt.Println("Hi...I am late")
	dataChan <- true
}