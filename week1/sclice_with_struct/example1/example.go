package main

import "fmt"



type Person struct {
	Fname string
	Lname string
	favFood []string

}

func (p Person)Walk() string{
	return fmt.Sprintln(p.Fname,"is walking")

}

func main() {
	p:=Person{"Shalini", "Singh" ,[]string{"Paneer","Subway","Biryani","Curd Rice"}}
	fmt.Println(p.Walk())
	fmt.Println(p.Fname)
	fmt.Println(p.favFood)

	for _, food := range p.favFood{
		fmt.Println(food)
	}

}