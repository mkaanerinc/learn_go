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

type displayer interface {
	Display()
}

// The `outputable` interface **embeds** both `saver` and `displayer` interfaces.
// This means any type that satisfies **both** `Save() error` and `Display()` methods,
// also satisfies the `outputable` interface.
//
// ✅ Interface embedding allows composing complex interfaces from simpler ones,
// making your code more modular and easier to reason about.
//
// For example, both `note.Note` and `todo.Todo` likely implement both `Save()` and `Display()`,
// so they are valid `outputable` types.
//
// This enables us to pass either of them to the `outputData()` function,
// which only accepts values implementing the `outputable` interface.
type outputable interface {
	saver
	displayer
}

func main() {
	printSomething(1)
	printSomething("Hello, World!")
	printSomething(3.14)
	printSomething(true)

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

	err = outputData(todo) // Save and display the todo item

	if err != nil {
		return
	}

	outputData(userNote) // Save and display the note
}

// printSomething takes an interface{} value and prints its type and value.
// It uses type assertions to determine the concrete type of the value.
// If the value is of type int, float64, or string, it prints the corresponding value.
// If the value is of any other type, it does nothing.
func printSomething(value interface{}) {
	// Use a type switch to determine the concrete type of the value
	intVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", intVal)
		return
	}

	float64Val, ok := value.(float64)

	if ok {
		fmt.Println("Float64:", float64Val)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println("String:", stringVal)
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

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func getNoteData() (title, content string) {
	title = getUserInput("Note title: ")
	content = getUserInput("Note content: ")

	return title, content
}
