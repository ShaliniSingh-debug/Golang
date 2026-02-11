package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type Truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t Truck)transportationDevice() string{
	return  fmt.Sprintln("i am a truck having ",t.vehicle.door,"doors ,and i have 4 wheels, its",t.fourWheel)

}

func (s sedan)transportationDevice() string{
	return  fmt.Sprintln("i am a sedan having ",s.vehicle.door,"doors ,i am luxury its",s.luxury)

}

type transpotation interface {
	transportationDevice() string
	
}

func report(trans transpotation){
	fmt.Println(trans.transportationDevice())

}

func main() {
	truck := Truck{vehicle{2, "red"}, true}
	car := sedan{vehicle{4, "Blue"}, false}

	report(truck)
	report(car)

}