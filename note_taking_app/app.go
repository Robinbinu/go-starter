package main

import (
	"fmt"
	"example.com/note/note"
)


func main(){
	newNote,_ := getNoteData()
	newNote.PrintNoteDetails()
}

func getNoteData()(newNote note.Note,err error){
	title:=getUserInput("Note title: ")
	content:=getUserInput("Note content: ")
	newNote,err= note.New(title,content)
	if err != nil{
		return note.Note{},err
	}
	return newNote,nil
}



func getUserInput(text string) (value string){
	fmt.Print(text)
	fmt.Scanln(&value)
	return value
}