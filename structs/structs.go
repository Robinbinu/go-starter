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

func newUser(firstName,lastname string) User{
	return User{
		firstName: firstName,
		lastName: lastname,
	}
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
	nUser := newUser("Bob","Joe")
	fmt.Println(nUser.firstName)
	fmt.Println(nUser.lastName)
	fmt.Println(nUser.timeCreated)
}


