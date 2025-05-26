package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}

/*
defer in Go
-----------
`defer` is a built-in keyword in Go that postpones the execution of a function or method
until the surrounding function (in which `defer` is declared) completes.

How it works:
-------------
- When a `defer` statement is encountered, Go schedules the deferred function to run **after the current function returns**.
- Deferred functions are executed in **LIFO (Last-In-First-Out)** order if there are multiple.

Common usage:
-------------
- Closing open files: `defer file.Close()`
- Unlocking mutexes: `defer mu.Unlock()`
- Recovering from panics: `defer recover()`

Benefits:
---------
- Ensures clean-up code is always run, even if an error occurs or a return happens early.
- Improves readability by placing cleanup near the resource allocation.
- Prevents resource leaks (open files, memory, network connections, etc.).

Code explanation:
-----------------
- `os.Open` tries to open a file. If it fails, an error is returned immediately.
- If the file is successfully opened, `defer file.Close()` ensures the file is closed **no matter what** happens afterward.
- `bufio.NewScanner` reads the file line by line.
- Each line is appended to a slice named `lines`.
- If a scanning error occurs, it's caught and reported.
- At the end, the lines are returned — and thanks to `defer`, the file is safely closed once the function ends.

Why `defer` here?
-----------------
If we didn't use `defer`, we would have to manually call `file.Close()` in every possible return path,
which is error-prone and harder to maintain.

Conclusion:
-----------
`defer` is ideal for resource cleanup. It simplifies code, avoids duplication, and guarantees execution.
*/
