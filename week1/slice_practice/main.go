package main

import "fmt"
/*
Initialize a SLICE of int using a composite literal;
print out the slice;
range over the slice
printing out just the index;
range over the slice printing out both the index and the value */

func main() {
	thisisSlice := []int{11,2,3,4,5,73,5,6,7,8}
	fmt.Println("Slice Practice",thisisSlice)

	for index := range thisisSlice{
		fmt.Println("index of the slice: ",index)
	}
	for index, value := range thisisSlice{
		fmt.Println("index of the slice: ",index ," ", "Value of the slice: ",value)
		// fmt.Println()
	}


/*
Initialize a MAP using a composite literal where the key is a string and the value is an int;
print out the map; range over the map printing out just the key; range over the map
printing out both the key and the value
*/

thisisMap:= make(map[string]int)

thisisMap["someone"] = 29
thisisMap["Sanjana"] = 30
thisisMap["Krishna"] = 35
thisisMap["Ram"] = 26

fmt.Println(thisisMap)

for index , value:= range thisisMap{
	fmt.Println(index , value)
}

for index := range thisisMap{
	fmt.Println(index)
}
}




