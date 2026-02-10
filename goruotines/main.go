package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go greet("its nice to meet you!",done)
	go greet("its great to see you again!",done)
	
	go slowGreet("how... are... you...??..",done)
	
	go greet("Thanks for coming!",done)
	for range done{
		// fmt.Println()
	}

}

/*
without concurrencry :
function executes after each other, second function will start once first function will finish
store("hello") ---> store("hi") ----> store("this is go routine")-----> and so on(one after another)


parallel execute concurrently :
store("hello")
store("hi")                    sum(2,3)
store("this is go routine")

*/

func greet(message string ,doneChan chan  bool ) {
	fmt.Println("Hi there i am shalini", message)
	doneChan <-true
}
func slowGreet(message string , doneChan chan  bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hi there i am shalini", message)
	doneChan <- true
	close(doneChan)
}
