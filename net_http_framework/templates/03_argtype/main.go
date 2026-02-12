package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
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

	nf , err := os.Create("Index.html")
	if err != nil{
		log.Fatal("failed to create the file " , err)

	}
	defer nf.Close()

	io.Copy(nf , strings.NewReader(str))
}