package main

import (
	"fmt"
	"strings"
	
)

func main() {
	firstName :=ReadString("Enter Your First Name : ")
	lastName :=ReadString("Enter Your First Name : ")
	email :=ReadString("Enter Your Email : ")
	age :=ReadInt("Enter Your age : ")

	user , err := NewUser(firstName , lastName , email ,int32(age))

	if err !=nil {
		fmt.Println("Error : " , err)
		return 
	}
	user.PrintDetails()

	// newEmail := ReadString("Enter a new Email to update old ")
	choice := strings.ToLower(ReadString("Do you want to update your email? (y/n): "))
	if choice == "y" || choice == "yes" {
		newEmail := ReadString("Enter new email: ")
		if err := user.UpdateEmail(newEmail); err != nil {
			println("Error:", err.Error())
		} else {
			println("Email updated successfully!")
		}

		// Print updated details
		user.PrintDetails()
	} else {
		println("Email not updated.")
	}
	
}


