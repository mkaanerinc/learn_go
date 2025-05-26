package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		//fmt.Println(doneChan)
	}
}

/*
Working with Multiple Channels
------------------------------
When using multiple goroutines that need to signal back when they finish,
you typically use **a separate channel for each goroutine**.
This allows you to track the completion of each one independently.

How It’s Done:
--------------
- You create a slice or array of channels (e.g., dones := make([]chan bool, 4)).
- Each goroutine gets its own channel.
- In the main function, you loop over the channels and wait for each to signal completion using `<-channel`.

Features:
---------
- Allows you to track which specific task is done.
- Prevents data races since channels aren’t shared across goroutines.
- Channels can be closed after sending to signal that no more data will come.

Select Statement (Advanced):
----------------------------
In more complex cases, Go provides the `select` statement to listen to multiple channels concurrently.
This is especially useful when you're not sure which one will return first.

Code Explanation:
-----------------
- The code defines two functions: `greet` and `slowGreet`, both send a `true` value to the given channel when done.
- A **single channel** named `done` is used and passed to all goroutines (this is problematic, explained below).
- `close(doneChan)` is called inside each function — this is **unsafe** when multiple goroutines share the same channel (only one goroutine should close it).
- In `main`, `for range done` tries to iterate over the channel, but once closed, it will stop — which can cause incomplete output.

Problem in Code:
----------------
- Using **one shared channel** for all goroutines and closing it in each is dangerous.
- Only one goroutine should close a channel, otherwise it causes panic.
- The correct approach is:
  1. Use **separate channels** for each goroutine.
  2. Wait on each channel (e.g., with a loop: `for _, ch := range dones { <-ch }`).

Recommendation:
---------------
Uncomment the `dones := make([]chan bool, 4)` part, assign a channel to each, and pass the corresponding channel to each goroutine.
Then use a loop to wait for each with `<-done`.

This ensures safe and correct handling of multiple channels in concurrent code.
*/
