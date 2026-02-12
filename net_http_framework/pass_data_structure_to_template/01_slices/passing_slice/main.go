package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template
func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))

}

func main() {
	Fruits :=[]string{"Apple","Orange","Grapes","Sweet Potato","berries","Strawberries"}
	err := tmpl.Execute(os.Stdout, Fruits)
	if err != nil{
		log.Fatal(err)
	}
	
}