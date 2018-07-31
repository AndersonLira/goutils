package io

import (
	"io/ioutil"
	"os"
)

//WriteFile writes a text file with content
//It returns error if some trouble happens
func WriteFile(fileName string, content string) error {
	b := []byte(content)
	err := ioutil.WriteFile(fileName, b, 0644)
	return err
}

//ReadFile reads a file and return its content as string
func ReadFile(fileName string) (content string, err error) {
	b, err := ioutil.ReadFile(fileName)
	return string(b), err
}

//AppendFile writes on file increasing its content
func AppendFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	defer file.Close()

	if err != nil {
		return err
	}

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
