// You can edit this code!
// Click here and start typing.
package main

import (
	// "errors"
	"fmt"
)

func main() {
	user := enterValue("Please Enter Your Name!")
	fmt.Println("hello", user)

	// remainder function
	largenumber := getNumbers("Please enter the First Number: ")
	smallnumber := getNumbers("Please enter the Second Number: ")

	final_output := getRemainders(largenumber, smallnumber)
	fmt.Println("the remainder is: ", final_output)

	// find even numbers
	values := findEvennumbers(100)
	fmt.Println(values)

	// fizzbuzz challenge
	fizzbuzz := listFizzBuzz(100)
	fmt.Print(fizzbuzz)

	// total numbers challenge
	outputs := getAddedNumber(1000)
	fmt.Print(outputs)

	// function as variable
	half := func(number int) (int, bool) {
		return number / 2, number%2 == 0
	}
	fmt.Println(half(1))
	fmt.Println(half(2))

	// largest number in the list
	greatest := max(1, 4, 6, 2, 34, 87, 6, 99, 67)
	fmt.Println("the largest number in the list is ", greatest)

	// fmt.Println((true && false) || (false && true) || !(false && false))

	// silly foo practice
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

}

func foo(number ...int) {
	fmt.Println(number)

}

/*Create a program that prints to the terminal asking for a user to enter
their name. Use fmt.Scan to read a user’s name entered at the
terminal. Print “Hello <NAME>” with <NAME> replaced with what the
user entered at the terminal*/

func enterValue(prompt string) string {
	var username string
	fmt.Println(prompt)
	fmt.Scan(&username)
	return username
}

/* the function ends here for the remainder function*/
