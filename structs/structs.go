package main

import(
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	timeCreated time.Time
}

func (user User) printDetails(){
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.timeCreated)
}

func (user User) clearName(){
	user.firstName = ""
	user.lastName = ""
}

func main(){
	newUser := User{
		firstName: "Joe",
		lastName: "Doe",
		timeCreated: time.Now(),
	}

	newUser.printDetails() //Joe Doe 2025-05-12 09:07:20.354820494 +0000 UTC m=+0.000034325
	newUser.clearName()
	newUser.printDetails() //Joe Doe 2025-05-12 09:07:20.354820494 +0000 UTC m=+0.000034325

}


