package main

import "fmt"

func main() {
	name := "Shalini Singh"

	templating := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
			<meta charset="UTF-8">
			<title>My Basic Page</title>
			
	</head>
	<body>
			<h1>` + name + `</h1>
	</body>
	</html>`
	fmt.Println(templating)

}