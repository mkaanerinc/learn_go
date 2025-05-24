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
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
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

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Note saved successfully.")
	return nil
}

func getNoteData() (title, content string) {
	title = getUserInput("Note title: ")
	content = getUserInput("Note content: ")

	return title, content
}

// The `saver` interface declares a contract:
// Any type that implements a method `Save() error` is considered a `saver`.
//
// Both `todo.Todo` and `note.Note` implement this interface by having a `Save` method,
// so they can be used interchangeably wherever a `saver` is expected.
//
// This allows for writing flexible and reusable code,
// such as a `saveData` function that accepts any type implementing `saver`:
//
//	func saveData(data saver) {
//	    err := data.Save()
//	    if err != nil {
//	        fmt.Println("Saving failed:", err)
//	    }
//	}
//
// ✅ Benefits of interfaces in Go:
// - Enable polymorphism (use different types the same way).
// - Help in testing (e.g., use mock implementations).
// - Encourage modular design and separation of concerns.
//
// ➕ Go interfaces are **implicit**, meaning types implement an interface just by having the required methods.
// There's no `implements` keyword like in Java or C#.
//
// For example:
// - If `note.Note` has a `Save() error` method, it **automatically** satisfies `saver`.
// - Same goes for `todo.Todo`.
