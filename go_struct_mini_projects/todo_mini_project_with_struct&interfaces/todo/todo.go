package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	
)

type Todo struct{

	Text string `json:"content"`
}

func NewNote(content string) (Todo , error) {
	if content ==""{
		return  Todo{}, errors.New("The value is Empty please check.")

	}
	return Todo{
		
		Text: content,
		
	}, nil

}


func (todo Todo) DisplayNote() {
	fmt.Printf("Your following text is:\n\n%v\n\n", todo.Text)


}
func (todo Todo) Save() error{
	fileName:= strings.ReplaceAll(todo.Text," ","_")
	fileName = "todo.json"
	json , err := json.Marshal(todo)

	if err !=nil{
		return err
	}
	return os.WriteFile(fileName,json,0664)
}

// func New(title , content string) (Note,error){
// 	if title =="" || content==""{
// 		return  Note{},errors.New("Invalid Input")
// 	}
// 	return Note{},nil
// }