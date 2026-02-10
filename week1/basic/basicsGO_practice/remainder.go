package main

import "fmt"


/*
Create a program that prints to the terminal asking for a user to enter
a small number and a larger number. Print the remainder of the larger
number divided by the smaller number. 
*/


func getNumbers(prompt string) (int){
	var number int
	fmt.Print(prompt)
	fmt.Scan(&number)
	return number
}

func getRemainders(large_num , small_num int) (int){
	var output int
	if large_num >small_num {
		output = large_num%small_num
		return output

	}else{
		output = small_num%large_num
		return output
	}
	
}