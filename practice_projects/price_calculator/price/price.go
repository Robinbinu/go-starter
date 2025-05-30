package price

import (
	"fmt"
	"pc/conversion"
	"pc/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager `json:"-"`
	Prices            []float64               `json:"prices"`
	TaxRate           float64                 `json:"tax_rate"`
	TaxIncludedPrices map[string]string       `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error{

	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}
	//convert string to float
	job.Prices, err = conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	return nil

}

func (job *TaxIncludedPriceJob) Process() error{
	err:=job.LoadData()
	if err != nil {
		return err
	}


	result := map[string]string{}

	for _, price := range job.Prices {
		formattedPrice := fmt.Sprintf("%.2f", price*(1+job.TaxRate))
		result[fmt.Sprintf("%.2f", price)] = formattedPrice
	}
	//did't work previously since pointer was not used so copy was created
	job.TaxIncludedPrices = result
	err = job.IOManager.WriteJson(job)
	if err != nil {
		return err
	}

	return nil
}

// constructor func specific to struct
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		Prices:    []float64{10, 20, 30},
		TaxRate:   taxRate,
		IOManager: fm,
	}
}
