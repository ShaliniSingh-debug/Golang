package main

import "fmt"

type gator int
type flamingo bool

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}
func (f flamingo) greeting() {
	fmt.Println("“Hello, I am pink and beautiful and wonderful.")

}

// interface
type swampCreature interface {
	greeting()
}

// interface passed to method you can use it directly bayou(value)==bayou(g1) from line 28
func bayou(s swampCreature) {
	s.greeting()

}

func main() {
	var g1 gator
	bayou(g1)
	g1 = 42
	// print value
	fmt.Println(g1)
	// prints type
	fmt.Printf("%T\n", g1)

	var x int
	x = int(g1)
	fmt.Println(x) // default value
	fmt.Printf("%T\n", x)

	// flamingo
	var f1 flamingo
	bayou(f1)

	f1 = true
	fmt.Println(f1) // default value
	fmt.Printf("%T\n", f1)

}
