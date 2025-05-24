package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	todo.Display()
	err = todo.Save()

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	fmt.Println("Todo saved successfully.")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Note saved successfully.")
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
