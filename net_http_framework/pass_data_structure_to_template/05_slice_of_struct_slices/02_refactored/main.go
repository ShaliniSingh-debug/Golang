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
	Manufacturer string
	Food string
	Type string
}

// type Diversity struct {
// 	States []Harvest
// 	Foodplnt []FoodPlant
// }

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
		Manufacturer: "Amul",
		Food: "Milk",
		Type: "Dairy",
	}
	Bakery := FoodPlant{
		Manufacturer: "Parle Products",
		Food: "Biscuits",
		Type: "Bakery",
	}

	Beverages := FoodPlant{
		Manufacturer: "Coca-Cola India",
		Food: "Frooti/Appy",
		Type: "Beverages",
	}
	foodtypes := []FoodPlant{Dairy ,Bakery , Beverages}

	harvests := []Harvest{cotton,oranges,Peanuts,rice,mango,coconut}
// two structs can be called like this aswell we dont need new struct to put the value of Harvest and FoodPlant
	items := struct{
    States   []Harvest
    Foodplnt []FoodPlant
	}{
			States:   harvests,
			Foodplnt: foodtypes,
	}


	err := tmpl.Execute(os.Stdout,items)
	if err != nil {
		log.Fatal(err)
	}



}