package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

func getUserInput(prompt string) (text string) {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin) // Create a new reader to handle long input
	// ReadString reads until the first occurrence of the delimiter (newline in this case)
	text, err := reader.ReadString('\n') // Read until newline

	if err != nil {
		return "" // Handle error by returning an empty string
	}

	text = strings.TrimSuffix(text, "\n") // Remove the newline character from the end
	text = strings.TrimSuffix(text, "\r") // Remove carriage return for Windows compatibility

	return text
}

func getNoteData() (title, content string) {
	title = getUserInput("Note title: ")
	content = getUserInput("Note content: ")

	return title, content
}
