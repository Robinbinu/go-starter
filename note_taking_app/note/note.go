package note

import (
	"errors"
	"fmt"
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

func (note Note) PrintNoteDetails(){
	fmt.Println(note.title)
	fmt.Println(note.content)
	fmt.Println(note.timeCreated)
}

