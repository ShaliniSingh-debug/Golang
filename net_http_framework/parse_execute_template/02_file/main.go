package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.ParseFiles("template.gohtml")
	if err != nil {
		log.Fatalln(err)

	}
	

	nf , err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file", err)
	}
	defer nf.Close()

	err = tmpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}


