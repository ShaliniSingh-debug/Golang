package func_variable_return

// import (
// 	"fmt"
// )

// this type is created to replicate func(int)int as TransFunc
type TranSFunc func(int) int

// pass function as parameter in function with normal value return not function return
func TransformNumbers(numbers *[]int, transform TranSFunc) []int {
	dNumbers := []int{}

	for _, value := range *numbers {
		dNumbers = append(dNumbers, transform(value))

	}
	return dNumbers
}

// return function and return value , not the function call
func TransformNumbersWithReturn(number *[]int) TranSFunc {
	if (*number)[0] == 1 {
		return DoubleNumbers
	} else {
		return TripleNumebr
	}
}

func DoubleNumbers(value int) int {
	return value * 2
}

func TripleNumebr(value int) int {
	return value * 3
}
