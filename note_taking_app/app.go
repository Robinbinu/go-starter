package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func errorFound(err error) bool{
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func main() {
	newNote, err := getNoteData()
	if errorFound(err){return}

	newNote.Display()

	err = newNote.Save()
	if errorFound(err){return}
	fmt.Println("note saved successfully")

}

func getNoteData() (newNote note.Note, err error) {
	title,_ := getUserInput("Note title: ")
	content,_ := getUserInput("Note content: ")
	newNote, err = note.New(title, content)
	if err != nil {
		return note.Note{}, err
	}
	return newNote, nil
}

func getUserInput(text string) (value string,err error) {
	fmt.Print(text)
	// fmt.Scanln(&value) commented because Scan often gets messy with strings separated
	reader := bufio.NewReader(os.Stdin)
	value,err = reader.ReadString('\n')
	if err != nil{
		return "",err
	}

	value = strings.TrimSuffix(value,"\n") //if not added the strings has a new line by default
	value = strings.TrimSuffix(value,"\r") //windows ends strings with \r\n

	return value,nil

}