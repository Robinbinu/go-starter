package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []float64{10,20,30}
	taxRates := []float64{0,0.07,0.1,0.15}

	result := map[float64][]float64{}

	for _,taxRate := range taxRates{
		for _,price := range prices{
			result[taxRate]=append(result[taxRate],calcTaxedPrice(taxRate,price))
		}
	}
	fmt.Print(result)
}

func calcTaxedPrice(taxRate float64,price float64) float64{
	price *= 1 + taxRate
	roundedPrice := math.Round(price*100)/100
	return roundedPrice
}