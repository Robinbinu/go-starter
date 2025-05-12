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

func main(){
	newUser := User{
		firstName: "Joe",
		lastName: "Doe",
		timeCreated: time.Now(),
	}

	newUser.printDetails();
}


