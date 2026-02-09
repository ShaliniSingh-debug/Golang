package main

import (
	"errors"
	"fmt"
	"time"
)


type Book struct{
	BookTitle string
	BookAuthor string
	BookPublishedYear int
	BookAvailable bool
	AddedDated time.Time
}

func NewBook(book , author string, publishYear int) (Book , error){
	if book =="" || author == "" || publishYear <=0{
		return Book{}  , errors.New("please check you input value something is missing , book not added")
		
	} 
	return Book{
		BookTitle: book,
		BookAuthor: author,
		BookPublishedYear: publishYear,
		BookAvailable: true,
		AddedDated:time.Now(),
	} , nil 
	

}


func AddBook(library *[]Book, book Book) {
	*library = append(*library, book)
	 
}
func ViewBook(library []Book){
	if len(library)==0{
		fmt.Println("No Book is available")
		return
	}

	fmt.Println("\n Book List")
	for i, book := range library{
		status :="Available"
		if !book.BookAvailable{
			status = "Borrowed"
		}

		fmt.Printf(
			"%d. %s | %s | %d | %s\n",
			i+1,
			book.BookTitle,
			book.BookAuthor,
			book.BookPublishedYear,
			status,
			
		)
	}
}

func BorrowBook(library []Book , index int) error{
	if index<0 || index>=len(library){
		return errors.New("Invalid Book Number")

	}

	if !library[index].BookAvailable{
		return errors.New("Book Already Borrowed")
	}
	library[index].BookAvailable=false
	return nil
}