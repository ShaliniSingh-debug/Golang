/*
create a struct that holds person fields
create a struct that holds secret agent fields and embeds person type
attach a method to person: pSpeak
attach a method to secret agent: saSpeak
create a variable of type person
create a variable of type secret agent
print a field from person
run pSpeak attached to the variable of type person
print a field from secret agent
run saSpeak attached to the variable of type secret agent
run pSpeak attached to the variable of type secret agent
*/
package main

import "fmt"

type Person struct {
	Name string
	lastName string
}

type SecretAgent struct {
	Person
	lkt bool
}

func (p Person) pSpeak() {
	fmt.Println(p.Name, `says , "Whatsup!"`)
}

func (sa SecretAgent) saSpeak() {
	fmt.Println(sa.Person.Name, `says , "has a license to Catch"`,sa.lkt)
}

func main() {
	// fmt.Println("Employee")
	p := Person{"Shalini" , "Singh"}
	sa :=SecretAgent{Person{"James" , "Bond"},true}
	fmt.Println(p.Name)
	p.pSpeak()
	fmt.Println(sa.Name)
	sa.saSpeak()
	sa.pSpeak()
}
