package price

import (
	"fmt"
	"pc/conversion"
	"pc/filemanager"
)

type TaxIncludedPriceJob struct{
	Prices []float64
	TaxRate float64
	TaxIncludedPrices map[string] string
}

func (job *TaxIncludedPriceJob) LoadData(){
	
	//read text from file
	path := "prices.txt"
	lines,err := filemanager.ReadLines(path)
	if err != nil {
		fmt.Println(err)
	}
	//convert string to float
	job.Prices,err = conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
	}


}

func (job *TaxIncludedPriceJob) Process(){
	job.LoadData()

	result := map[string]string{}

	for _,price := range job.Prices{
		formattedPrice := fmt.Sprintf("%.2f",price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f",price)] = formattedPrice
	}
	//did't work previously since pointer was not used so copy was created
	job.TaxIncludedPrices = result
	err:=filemanager.WriteJson(fmt.Sprintf("result_taxRate_%.2f.json",job.TaxRate),job)
	if err != nil {
		fmt.Println(err)
	}
}

//constructor func specific to struct
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		Prices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}


