package main




// find the even numbers from 0 to 100

func findEvennumbers(number int) []int {
	var evens []int
	for i := 0; i <= number; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}
	return evens
	
}