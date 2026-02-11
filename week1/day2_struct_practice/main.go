package main

import (
	// "fmt"
	"practice_struct/area"
)


func main(){

	// circle function starts here
	
	circle_output := area.Circle{2}

	
	area.Info(circle_output)

	// square function starts here
	sqaure := area.Square{2.5}
	
	
	area.Info(sqaure)
}