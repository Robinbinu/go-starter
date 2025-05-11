package main

import (
	"fmt"
)

func main1(){
	var revenue,expenses,tr,EBT,profit,ratio float64

	revenue = getInput("Enter revenue: ")
	expenses = getInput("Enter expenses: ")
	tr = getInput("Enter tax rate: ")

	EBT,profit,ratio = calculateVals(revenue,expenses,tr)

	fmt.Println(EBT)
	fmt.Println(profit)
	fmt.Println(ratio)

}


func getInput(text string) (userInput float64){
	fmt.Print(text)
	fmt.Scan(&userInput)
	return
}

func calculateVals(revenue,expenses,tr float64)(ebt,profit,ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1 - tr/100)
	ratio = profit/ebt
	return
}

// change