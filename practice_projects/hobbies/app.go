package main

import "fmt"

func main(){
	hobbies := []string{"eating","coding","sleeping"}
	fmt.Println(hobbies)
	// for i:=0;i<len(hobbies);i++{
	// 	fmt.Println(hobbies[i])
	// }
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	slc1 := hobbies[0:2]
	slc1alt := hobbies[:2]
	fmt.Println(slc1)
	fmt.Println(slc1alt)
	slc1alt = append(slc1alt, hobbies[1],hobbies[len(hobbies)-1])
	fmt.Println(slc1alt)


}