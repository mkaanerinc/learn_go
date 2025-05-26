package main

import (
	"fmt"
	"time"
)

// func greet(phrase string) {
// 	fmt.Println("Hello!", phrase)
// }

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {
	// go greet("Nice to meet you!")
	// go greet("How are you?")
	done := make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!")
	<-done
}

/*
What are Go Routines?
---------------------
Go Routines are lightweight threads managed by the Go runtime.
They enable concurrent execution of functions.
A function is turned into a Go Routine by prefixing it with the "go" keyword.
Example: go greet("Hello") → runs this function concurrently without blocking the main flow.

Why are They Used?
------------------
- To run multiple tasks at the same time (e.g., I/O operations, network calls)
- To improve performance
- To execute tasks in the background without blocking the main thread

Features:
---------
- Lightweight: Thousands of Go routines can be started.
- Small initial stack size (~2KB), grows dynamically as needed.
- Managed by the Go runtime.
- Enables concurrency (and parallelism with proper settings like GOMAXPROCS).

What is a Channel?
------------------
A Channel is used to communicate between Go Routines.
Defined using the "chan" keyword.
Send data: chan <- value
Receive data: value := <-chan

Why are Channels Used?
----------------------
- To safely share data between Go Routines
- For synchronization — e.g., waiting for one task to complete
- To avoid issues like deadlocks and race conditions

Code Explanation:
-----------------
- The `slowGreet` function simulates a slow task by sleeping for 3 seconds before printing.
- It accepts a channel (`doneChan`) and sends `true` into it once the task is done.
- In `main`, a channel is created: `done := make(chan bool)`.
- `slowGreet` is called as a Go Routine.
- `main` waits for the signal from the `done` channel using `<-done`.
- This ensures the program doesn’t exit before the Go Routine finishes.
*/
