package prices

import (
	"errors"
	"fmt"

	"tax_calculator.com/conversion"
	"tax_calculator.com/filemanager"
)

// tax Struct
type TaxIncludedPricesJob struct{
	IOManager filemanager.FileManager
	TaxRate float64
	InputPrices []float64
	TaxIncludingPrices map[string]string
}


// function that includes Struct call
func (job *TaxIncludedPricesJob ) readData()error{
	lines , err := job.IOManager.ReadFiles()
	if err!=nil{
		fmt.Println(err)
		return errors.New("DId not find the Files")
	}
	prices, err :=conversion.StringToFloatConversion(lines)

	if err != nil{
		fmt.Println(err)
		return errors.New("DId not find the Files to change")
	}
	job.InputPrices=prices
	return  nil
	
	}



// this also includes Struct 
func (job *TaxIncludedPricesJob ) Process() error {

	err := job.readData()
	if err != nil{
		return err
	}
	result := make(map[string]string)
	
	for _, price := range job.InputPrices {
		taxIncludeprice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxIncludeprice)

	}
	
	job.TaxIncludingPrices = result

	/* here we are writing and saving the file */ 

	job.IOManager.WriteJSONFile(job)
	return nil
		
}

// this is specially for struct because it starts with New it returns struct aswell
func NewTaxIncludedPriceJob(fm filemanager.FileManager,  taxRate float64) *TaxIncludedPricesJob {
	return &TaxIncludedPricesJob{
		InputPrices: []float64{10.29, 20.3, 43.89},
		TaxRate: taxRate,
		IOManager: fm,
	}
}

  