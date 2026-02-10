package main

import (
	"fmt"
	"tax_calculator.com/filemanager"
	"tax_calculator.com/prices"
)

func main() {
	// var prices []float64 = []float64{10.29, 20.3, 43.89}
	var taxRate []float64 = []float64{0, 0.07, 0.1, 0.15}
	

	for _, taxRate := range taxRate {
		fm := filemanager.New("prices.txt" , fmt.Sprintf("result_%.0f.json",taxRate*100) )
		pricejob := prices.NewTaxIncludedPriceJob(fm ,taxRate)
		err :=pricejob.Process()
		if err !=nil{
			fmt.Println("Could not process the job")
			fmt.Println(err)
		}

		
	}


}
