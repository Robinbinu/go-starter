package main

import "fmt"

func main() {
	alphabets := map[string]string{
		"B": "Bat",
		"C": "Cat",
	}
	alphabets["A"] = "Apple"
	fmt.Println(alphabets) //map[A:Apple B:Bat C:Cat]

	delete(alphabets,"C")
	fmt.Println(alphabets) //map[A:Apple B:Bat]

	
}
