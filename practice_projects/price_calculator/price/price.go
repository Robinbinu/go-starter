package price

import "fmt"

type TaxIncludedPriceJob struct{
	Prices []float64
	TaxRate float64
	TaxIncludedPrices map[string] float64
}

func (job TaxIncludedPriceJob) Process(){
	result := map[string]float64{}
	for _,price := range job.Prices{
		result[fmt.Sprintf("%.2f",price)] = price*(1+job.TaxRate)
	}
	fmt.Println(result)
}

//constructor func specific to struct
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		Prices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}

