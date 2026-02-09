package main

import (
	"errors"
	"fmt"
	"time"
)



type UserDetails struct{
	UserFirstName string
	UserLastName string
	UserEmail string
	UserAge int32
	CreatedAt time.Time
}

func NewUser(firstName string, lastName string, userEmail string, userAge int32) (*UserDetails ,error){
	if userAge<=0 || userEmail=="" || firstName=="" ||lastName==""{
		return nil , errors.New("User Details are invalid please check and enter corrrectly ")

	} 
	return &UserDetails{
	UserFirstName: firstName,
	UserLastName: lastName,
	UserEmail: userEmail,
	UserAge: userAge,
	CreatedAt: time.Now(),
} , nil

}

// give users Full name 
func (u *UserDetails) Fullname() string{
	return u.UserFirstName + " "+u.UserLastName
}

// check if user is Adult
func (u *UserDetails) IsAdult() bool{
	return u.UserAge >=18
			
}



//update email of the user
func (u *UserDetails) UpdateEmail(newEmail string) error{
	if newEmail == ""{
		return errors.New("the email is required")

	}
	u.UserEmail = newEmail
	return nil
}

// print details of the user

func (u *UserDetails) PrintDetails() {
	fmt.Println("--------- User Details ---------")
	fmt.Println("Full Name:", u.Fullname())
	fmt.Println("Email:", u.UserEmail)
	fmt.Println("Age:", u.UserAge)
	fmt.Println("Adult:", u.IsAdult())
	fmt.Println("Created At:", u.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("--------------------------------")
}