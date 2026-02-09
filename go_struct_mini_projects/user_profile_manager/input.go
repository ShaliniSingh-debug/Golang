package main

import (
	"fmt"
)

func ReadString(prompt string) string{
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func ReadInt(prompt string) int {
	var input int
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
} 	