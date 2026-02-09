package main



func max(number ...int) int {
	var largest int
	for _,v := range number {
		if v>largest{
			largest = v
		}
	}
	return largest
}