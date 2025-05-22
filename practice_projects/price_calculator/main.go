package main

import (
	"fmt"
	"pc/filemanager"
	price "pc/price"
)

func main() {
	taxRates := []float64{0,0.07,0.1,0.15}
	for _,taxRate := range taxRates{
		filemanager := filemanager.New("prices.txt",fmt.Sprintf("result_taxRate_%.2f.json",taxRate))
		job:=price.NewTaxIncludedPriceJob(filemanager,taxRate)
		job.Process()
		fmt.Println(job)
	}

	

}
