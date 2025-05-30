package main

import "fmt"

func main(){
    var arr = []int{1,2,3,4} //alt decl: var arr = [4]int{}
    n := len(arr) //len of array - no of elements present
    c := cap(arr) //capacity the array can hold
    strArr := []string{"str1","str2","str3"}

   fmt.Println(arr) //[1 2 3 4]
   fmt.Println(strArr[1]) //str2

   slicedArr := arr[1:3] //here [a:b] elements from a->b-1 are sliced
   slicedArr1 := arr[:3] //[a:]

   fmt.Print(slicedArr)
   fmt.Println(slicedArr1)
   fmt.Println(n)
   fmt.Println(c)
   toBeAppended := []int{1,3,4}
   apArr := append(arr,toBeAppended...) // . . . used to unpack the list
   fmt.Println(apArr)
}