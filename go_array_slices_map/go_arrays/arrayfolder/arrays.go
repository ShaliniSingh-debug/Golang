package main

import 
("fmt"
)


type Product struct{
	id string
	title string
	price float64
}

func main() {
//  add list of prices in list of prices 
	prices := []float64{30.33,56.67,98.00}
	newPrices :=[]float64{20.33,56.89,78.45}
	prices = append(prices, newPrices...)
	fmt.Println(prices)

	// task 1
	hobbies := [3]string{"Playing" ,"jumping","roaming"}
	fmt.Println(hobbies)

	// task 2 

	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// task 3

	namehobbies := hobbies[:2]
	fmt.Println(namehobbies)

	// task 4
	fmt.Println(cap(namehobbies))
	namehobbies =namehobbies[1:3]
	fmt.Println(namehobbies)


	// task 5

	courseGoals := []string{"Learn Go","Learn it fast"}
	fmt.Println(courseGoals)


	// task 6

	courseGoals[1]="Learn All Details"
	courseGoals = append(courseGoals, "Leanr all the Basics Carefully!")

	fmt.Println(courseGoals)

	// task 7

	products := []Product{
	{
		"first",
		"A first Product",
		20.57,
	 } ,
	 {
		"second",
		"A second Product" ,
		 20.33,
		},
	}
	// fmt.Println(products)


	newProduct := Product{
		"third",
		"This is Third Product",
		30.33,
	}
	products = append(products, newProduct)
	fmt.Println(products)
}