package note

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct{
	title string
	content string
	timeCreated time.Time
}

func New(title,content string)(Note,error){
	if title == "" || content == ""{
		return Note{},errors.New("invalid input")
	}

	return Note{
		title: title,
		content: content,
		timeCreated: time.Now(),
	},nil
}

func (note Note) Display(){
	fmt.Printf("Title: %v\nContent: %v\nCreated on:%v\n",note.title,note.content,note.timeCreated)
}

func (note Note) Save() error{
	noteText := fmt.Sprintf("%v\n%v\nCreated on: %v",note.title,note.content,note.timeCreated)
	return os.WriteFile("Note.txt",[]byte(noteText),0644)
}

