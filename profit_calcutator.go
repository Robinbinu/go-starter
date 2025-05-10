package main

import "fmt"

func main(){
	var revenue,expenses,tr,EBT,profit,ratio float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&tr)

	EBT = revenue - expenses
	profit = EBT - (EBT*tr/100)
	ratio = profit/EBT

	fmt.Printf("EBT: %f",EBT)
	fmt.Printf("Profit: %f",profit)
	fmt.Printf("Ratio: %f",ratio)


}