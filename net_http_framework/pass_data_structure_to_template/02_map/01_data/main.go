package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template
func init() {
	tmpl = template.Must(template.ParseFiles("template.gohtml"))

}

func main() {
	// fruits := []string{"Apple","Orange","BlueBerries","Blackberries"}
	Harvest :=map[string]string{
		"Uttar Pradesh" :"Black Berries",
		"Madhya Pradesh" :"Guava",
		"Telangana": "Cotton",
		"Assam" :"Rice",
		"Keral":"Coconut",
		"Karnataka": "Peanuts",
	}

	 err := tmpl.Execute(os.Stdout, Harvest)
	 if err !=nil {
		log.Fatal(err)
	 }
}