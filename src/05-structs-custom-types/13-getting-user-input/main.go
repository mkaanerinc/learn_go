package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Note title: %s\nNote content: %s\n", title, content)
}

func getUserInput(prompt string) (value string, err error) {
	fmt.Print(prompt)
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil
}

func getNoteData() (title, content string, err error) {
	title, errorForTitle := getUserInput("Note title: ")

	if errorForTitle != nil {
		return "", "", errorForTitle
	}

	content, errorForContent := getUserInput("Note content: ")

	if errorForContent != nil {
		return "", "", errorForContent
	}

	return title, content, nil
}
