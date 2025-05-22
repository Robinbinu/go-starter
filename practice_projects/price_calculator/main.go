package main

import (
	"fmt"
	price "pc/price"
)

func main() {
	taxRates := []float64{0,0.07,0.1,0.15}
	for _,taxRate := range taxRates{
		job:=price.NewTaxIncludedPriceJob(taxRate)
		job.Process()
		fmt.Println(job)
	}

	

}
