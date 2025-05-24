package main

import (
	"fmt"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote.Display()
}

func getUserInput(prompt string) (value string) {
	fmt.Print(prompt)
	fmt.Scanln(&value)

	return value
}

func getNoteData() (title, content string) {
	title = getUserInput("Note title: ")
	content = getUserInput("Note content: ")

	return title, content
}
