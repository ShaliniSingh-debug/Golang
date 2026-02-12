package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Shalini Singh"

	str := fmt.Sprint(
		`
		<!DOCTYPE html>
	<html lang="en">
	<head>
			<meta charset="UTF-8">
			<title>My Basic Page</title>
			
	</head>
	<body>
			<h1>` + name + `</h1>
	</body>
	</html>`) 
	nf , err :=os.Create("index.html")
	if err !=nil{
		log.Fatal("Error creating the file ", err)
	}
	defer nf.Close()
	io.Copy(nf , strings.NewReader(str))
}