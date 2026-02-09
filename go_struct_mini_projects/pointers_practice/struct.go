package main

import (
	// "errors"
	"fmt"
	// "time"
	"user_details/user"
)



func main() {
	userFirstName := getUserData("Please Enter your first name: ")
	userLastName := getUserData("Please Enter your Last name: ")
	userBirthDate := getUserData("Please Enter your birth date (MM/DD/YYYY): ")
	
	var appUser *user.User

	// appUser , err := newUser(userFirstName , userLastName , userBirthDate)

	appUser , err := user.NewUser(userFirstName,userLastName,userBirthDate)
	if err != nil{
		fmt.Println(err)
		return
	}
	appUser.OutputUserDetails()
	// outputUserDetails(&appUser) you can user this values through pointers aswell

}

// func outputUserDetails(userDetail *user) { this is how you will access the pointer 

// func (u *user) outputUserDetails() {
// 	fmt.Println(u.firstName,u.lastName,u.birthDate)
// }

func getUserData(promtTxt string) string {
	fmt.Print(promtTxt)
	var value string
	fmt.Scanln(&value)
	return value
}

// func (u *User) outputUserDetails() {
// 	fmt.Println(u.firstName,u.lastName,u.birthDate)
// }

