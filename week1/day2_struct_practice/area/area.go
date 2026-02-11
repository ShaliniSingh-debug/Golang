package area

// import "math"
import "fmt"

const PI = 3.14

type Circle struct {
	Radius float64
}
type Square struct {
	Side float64
}

type Shape interface{
	GetArea() float64
}

// func NewCircle(radius float64) Circle {
// 	return Circle{
// 		Radius: radius,
// 	}

// }


func (c Circle) GetArea() float64 {
	area := PI * c.Radius * c.Radius
	// println("Area of Circle is : " , area)
	return area
}


// func NewSquare(side float64) Square{
// 	return Square{
// 		Side: side,
// 	}
// }

func (s Square) GetArea() float64{
	output := s.Side*s.Side
	// println("The Area of Square is: ",output)
	return output

}

func Info(s Shape) {
	fmt.Println(s.GetArea())
}


