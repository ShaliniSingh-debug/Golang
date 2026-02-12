package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template
type Harvest struct {
	States string
	Yeilds string
}

func init(){
	tmpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	Coconut := Harvest{
		States: "Kelara",
		Yeilds: "Coconuts",
	}

	err := tmpl.Execute(os.Stdout , Coconut)
	if err != nil {
		log.Fatal(err)
	}

}