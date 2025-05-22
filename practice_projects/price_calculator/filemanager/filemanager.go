package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

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
func WriteJson(path string, data any) error {
	file, err := os.Create(path)
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
