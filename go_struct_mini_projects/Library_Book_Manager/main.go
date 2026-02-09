package main

import "fmt"

func main() {
	var library []Book

	for {
		fmt.Println("\n📚 Library Book Manager")
		fmt.Println("1. Add Book")
		fmt.Println("2. View Books")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			book, err := getBookInput()
			if err != nil {
				fmt.Println("❌", err)
				continue
			}
			AddBook(&library, book)
			fmt.Println("✅ Book added")

		case 2:
			ViewBook(library)

		case 3:
			fmt.Print("Enter book number: ")
			var index int
			fmt.Scan(&index)

			err := BorrowBook(library, index-1)
			if err != nil {
				fmt.Println("❌", err)
			} else {
				fmt.Println("📕 Book borrowed successfully")
			}

		case 4:
			fmt.Println("👋 Goodbye")
			return

		default:
			fmt.Println("❌ Invalid option")
		}
	}
}

func getBookInput() (Book, error) {
	var title, author string
	var year int

	fmt.Print("Enter Title: ")
	fmt.Scan(&title)

	fmt.Print("Enter Author: ")
	fmt.Scan(&author)

	fmt.Print("Enter Publish Year: ")
	fmt.Scan(&year)

	return NewBook(title, author, year)
}
