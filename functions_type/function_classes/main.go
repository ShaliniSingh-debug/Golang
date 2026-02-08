package main

import (
	"fmt"

	"function_types/anonymus" // this are the module imports from other directories
	fvr "function_types/func_variable_return"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 3, 4, 1}

	// anonymous function
	doubled := anonymus.TransformNumbers(&numbers, func(n int) int {
		return n * 2
	})

	fact := anonymus.Factorial(5)

	// function returned as variable
	fn1 := fvr.TransformNumbersWithReturn(&numbers)
	fn2 := fvr.TransformNumbersWithReturn(&moreNumbers)

	result1 := fvr.TransformNumbers(&numbers, fn1)
	result2 := fvr.TransformNumbers(&moreNumbers, fn2)

	fmt.Println(doubled)
	fmt.Println(fact)
	fmt.Println(result1)
	fmt.Println(result2)
}
