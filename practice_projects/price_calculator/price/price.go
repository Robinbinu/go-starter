package price

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct{
	Prices []float64
	TaxRate float64
	TaxIncludedPrices map[string] float64
}

func (job *TaxIncludedPriceJob) LoadData(){
	file,err := os.Open("prices.txt")
	if(err != nil){
		fmt.Println("File Open Failed!!!")
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Reading file failed!")
		fmt.Println(err)
	}

	//convert string to float
	result := make([]float64,len(lines))

	for lineIndex,line := range lines{
		result[lineIndex],err = strconv.ParseFloat(line,64)
		if err != nil {
			fmt.Println("Error Converting String to float")
			fmt.Println(err)
		}
	}

	job.Prices = result


}

func (job *TaxIncludedPriceJob) Process(){
	result := map[string]float64{}
	for _,price := range job.Prices{
		result[fmt.Sprintf("%.2f",price)] = price*(1+job.TaxRate)
	}
	//did't work previously since pointer was not used so copy was created
	job.TaxIncludedPrices = result
}

//constructor func specific to struct
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		Prices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}


