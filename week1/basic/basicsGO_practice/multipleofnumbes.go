package main



func getAddedNumber(n int) int{
	var output int
	for i:= 0 ; i <n ; i++{
		if i%3 ==0 || i%5==0 {
			output += i

		}
	}
	return output
}