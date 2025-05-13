package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"` //struct tag, used to let json package know to use this alias as the key
	Content     string 	  `json:"content"`
	TimeCreated time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:       title,
		Content:     content,
		TimeCreated: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Title: %v\nContent: %v\nCreated on: %v\n", note.Title, note.Content, note.TimeCreated)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"

	jsonData, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}
