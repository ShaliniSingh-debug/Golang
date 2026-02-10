package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)



type FileManager struct {
	InputFile string
	OutputFile string
}
/*reading data from .txt files as i am passing path of the file as string and 
returing the string and error as slice of string (the data will go like array form)([]string , error) together , closing the file 



*/

func (fm FileManager )ReadFiles() ([]string , error){
	file , err := os.Open(fm.InputFile)

	if err !=nil{
		fmt.Println("Could not open the file!")
		fmt.Println(err)
		return nil , errors.New("failed to open the file!")
	
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil{
		fmt.Println("error Reading the file content")
		fmt.Println(err)
		file.Close()
		return nil , errors.New("Failed to scan the Data!")
	}
	file.Close()
	return lines , nil

}


/* write JSON data to FIles and saving the file with its rates basis
 each file will be created what tax rate is given eg : if taxrate is 0 then file will me created result_0.json
 if there is slice of rates the the files will also me many as much as length of slice  */


func (fm FileManager)WriteJSONFile(data interface{}) error{
	file , err := os.Create(fm.OutputFile)
	if err !=nil{
		return errors.New("Failed to create file.")
	}

	time.Sleep(3*time.Second)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err !=nil{
		file.Close()
		return  errors.New("Failed to convert data to Json.")
	}
	file.Close()
	return nil
}


/* here we have created a method for FileManager struct where the data is passed as inputfile and output
*file now why we did this .
*we did this to pass this value to other function because we cant pass direct struct into another functions and call struct directly
*again this New will be call to get this data as inputFle and OutputFile from anywhere like 

*filemanager.New("prices.txt" , fmt.Sprintf("result_%.0f.json",taxRate*100) )
here you can see input and output value is passed

*now when we want to use this  input file  this struct will work as receiver in that function

*the receiver we write as (fm FileManager -> this is the struct call ) the value from
new function will be receive in any function which uses this receiver 
eg. func (fm FileManager)WriteJSONFile(data interface{}) error{ 




*/ 
func New(inputFile , outputFile string) FileManager{
	return FileManager{
		InputFile: inputFile,
		OutputFile: outputFile,
	}
}