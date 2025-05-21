package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Content string `json:"content"`
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Content: content,
	}, nil
}

func (todo Todo) Display() {
	fmt.Printf("Content: %v\n", todo.Content)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	jsonData, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}
