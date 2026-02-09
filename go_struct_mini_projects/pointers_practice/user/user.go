package user

import (
	"fmt"
	"errors"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// this is constructure like oops concept for struct 
func NewUser(firstName , lastName, birthDate string)(*User,error){
	if firstName =="" ||lastName =="" || birthDate ==""{
		return nil , errors.New("first name , last name , birthdate is missing and required for the fields to operate")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName,u.lastName,u.birthDate)
}

