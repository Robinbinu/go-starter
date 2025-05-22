package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct{
	FileInputPath string
	FileOutputPath string
}

func New(fileInputPath string,fileOutputPath string) FileManager{
	return FileManager{
		FileInputPath: fileInputPath,
		FileOutputPath: fileOutputPath,
			
	}
}


func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.FileInputPath)

	if err != nil {
		return nil, errors.New("file open failed")
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("read lines failed")
	}

	file.Close()
	return lines, nil

}

// accepts any type of data, `any` is alt for `interface:{}`
func (fm FileManager)WriteJson(data any) error {
	file, err := os.Create(fm.FileOutputPath)
	if err != nil {
		return errors.New("error opening file")
	}

	encoder:=json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("error while encoding json")
	}

	return nil
}
