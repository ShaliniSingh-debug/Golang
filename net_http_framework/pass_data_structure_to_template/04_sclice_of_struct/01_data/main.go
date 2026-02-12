package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

type Harvest struct{
	State string
	Yeild string
}

func init(){
	tmpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main(){
	cotton :=Harvest{
		State: "Telangana",
		Yeild: "Cotton",
	}
	coconut :=Harvest{
		State: "Kelara",
		Yeild: "Cotton",
	}
	mango :=Harvest{
		State: "Uttar Pradesh",
		Yeild: "Mango",
	}
	rice :=Harvest{
		State: "Assam",
		Yeild: "rice",
	}
	Peanuts :=Harvest{
		State: "karnataka",
		Yeild: "Peanuts",
	}
	oranges :=Harvest{
		State: "Maharastra",
		Yeild: "oranges",
	}

	harvest := []Harvest{cotton,oranges,Peanuts,rice,mango,coconut}

	err := tmpl.Execute(os.Stdout , harvest)
	if err != nil{
		log.Fatal(err)
	}

}