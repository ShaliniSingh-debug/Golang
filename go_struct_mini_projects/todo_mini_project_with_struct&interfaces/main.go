package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "errors"
	"cunstructer_deep_dive/todo"
)
type saver interface {
	Save() error
}
type outputable interface {
	saver
	DisplayNote()

}

func main(){
	title , content :=getNoteData()
	todoText := getUserInput("Todo Text:")

	todo , err := todo.NewNote(todoText)
	if err !=nil {
		fmt.Println(err)
		return
	}


	userNote , err :=NewNote(title , content )
	if err != nil{
		fmt.Println(err)
		return
	}

	
	
	err =outputData(todo)
	if err != nil{
		return 
	}
	


	outputData(userNote)
	

}


func outputData(data outputable) error{
	data.DisplayNote()
	return data.Save()

}

func dataSaver( data saver)  error {
	err := data.Save()
	if err !=nil{
		fmt.Println("Saving the note Failed.")
		return err

	}
	fmt.Println("Saving the note Succeeded!")
	return nil

}

// get the  title and content
func getNoteData()(string , string ){
	title := getUserInput("please enter the Note title: ")
	
	content := getUserInput("please write you Note content here: ")
	
	return title , content 
}

func getUserInput(prompt string) (string){
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text , err := reader.ReadString('\n')
	if err !=nil{
		return ""
	}
	text =strings.TrimSuffix(text , "\n")
	text =strings.TrimSuffix(text , "\r")
	
	return text

}