package main

import (
	"log"
	"os"
	"text/template"
)

var tmpl *template.Template

type Harvest struct {
	State string
	Yeild string
}

type FoodPlant struct {
	Manufaturer string
	Food string
	Type string
}

type Diversity struct {
	States []Harvest
	Food []FoodPlant
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

	Dairy := FoodPlant{
		Manufaturer: "Amul",
		Food: "Milk",
		Type: "Dairy",
	}
	Bakery := FoodPlant{
		Manufaturer: "Parle Products",
		Food: "Biscuits",
		Type: "Bakery",
	}

	Beverages := FoodPlant{
		Manufaturer: "Coca-Cola India",
		Food: "Frooti/Appy",
		Type: "Beverages",
	}
	foods := []FoodPlant{Dairy ,Bakery , Beverages}

	states := []Harvest{cotton,oranges,Peanuts,rice,mango,coconut}

	items :=Diversity{
		States: states,
		Food: foods,
	}

	err := tmpl.Execute(os.Stdout,items)
	if err != nil {
		log.Fatal(err)
	}



}