package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputable interface{
	Display()
	saver //embedded interface
}

func errorFound(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func main() {
	newNote, err := getNoteData()
	if errorFound(err) {
		return
	}

	newTodo, err := getTodoData()
	if errorFound(err) {
		return
	}

	err = outputFile(newNote)
	if errorFound(err){return}
	fmt.Println("note saved successfully")

	outputFile(newTodo)
	fmt.Println("Todo saved sussesfully")

}

func getNoteData() (newNote note.Note, err error) {
	title, _ := getUserInput("Note title: ")
	content, _ := getUserInput("Note content: ")
	newNote, err = note.New(title, content)
	if errorFound(err) {
		return note.Note{}, err
	}
	return newNote, nil
}

func getTodoData() (newTodo todo.Todo, err error) {
	content, _ := getUserInput("Enter Todo: ")
	newTodo, err = todo.New(content)
	if errorFound(err) {
		return todo.Todo{}, err
	}
	return newTodo, nil
}

func getUserInput(text string) (value string, err error) {
	fmt.Print(text)
	// fmt.Scanln(&value) commented because Scan often gets messy with strings separated
	reader := bufio.NewReader(os.Stdin)
	value, err = reader.ReadString('\n')
	if errorFound(err) {
		return "", err
	}

	value = strings.TrimSuffix(value, "\n") //if not added the strings has a new line by default
	value = strings.TrimSuffix(value, "\r") //windows ends strings with \r\n

	return value, nil
}

func saveFile(data saver) (err error) { //function implements the saver interface
	err = data.Save()
	if errorFound(err) {
		return
	}
	return nil
}

func outputFile(data outputable)(err error){
	data.Display()
	err = saveFile(data)
	if errorFound(err){return err}
	return nil
}
