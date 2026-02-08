package anonymus

// import (
// 	"fmt"
// )

type TranSFunc func(int) int

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	// moreNumbers := []int{5,3, 4, 1}
// 	doubled := tranformNumbers(&numbers , func (number int) int{
// 		return number*2
// 	})
// 	// tripled := tranformNumbers(&numbers , tripleNumebr)

// 	// fmt.Println(doubled)
// 	// fmt.Println(tripled)

// 	// transformFn1 := tranformNumberswithReturn(&numbers)
// 	// transformFn2 := tranformNumberswithReturn(&moreNumbers)

// 	// tranFormedNumber :=tranformNumbers(&numbers,transformFn1)
// 	// tranFormedmoreNumber :=tranformNumbers(&moreNumbers,transformFn2)



// 	fact := factorial(5)
// 	fmt.Println(doubled)
// 	fmt.Println(fact)
// 	// fmt.Println(tranFormedmoreNumber)

// }

// pass function as parameter in function with normal value return not function return
func TransformNumbers(numbers *[]int , transform TranSFunc) []int{
	dNumbers := []int{}

	for _ ,value :=range *numbers{
		dNumbers = append(dNumbers, transform(value))

	}
	return dNumbers
}



func Factorial(number int) int{
	if number ==0{
		return 1
	}
	return number * Factorial(number-1)
}