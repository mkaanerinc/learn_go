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
Managing Channels with `select`
-------------------------------
In Go, the `select` statement is used to wait on multiple channel operations at once.
It's like a `switch`, but for channels. Only one of the ready channels will proceed at a time.

What It Does:
-------------
- Listens to multiple channels at once
- Executes the first case whose channel is ready
- Prevents blocking if one of the channels is already ready
- Supports `default` to avoid blocking when no channel is ready

Key Features:
-------------
- Only one `case` is selected at a time (if multiple are ready, one is picked randomly)
- Ideal for handling multiple possible outcomes (e.g., success vs. error)
- Helps simplify code when dealing with multiple goroutines and channels

Code Explanation:
-----------------
- The code defines tax rates and creates a `doneChans` and `errorChans` slice for each tax rate.
- For each tax rate:
  - A `filemanager` and `priceJob` are created.
  - A new goroutine is started with `Process`, which takes both done and error channels.
- The second loop iterates over each job's channels.
  - The `select` statement waits for **either**:
    - A message on the `error` channel → if there's an error, it's printed.
    - A signal on the `done` channel → prints "Done!" when the job finishes.
- This pattern ensures:
  - The main function doesn't block forever.
  - Both success and failure paths are handled concurrently and cleanly.

Why It’s Useful:
----------------
- Handles both success and failure gracefully.
- Avoids deadlocks by ensuring the main goroutine can respond to whichever channel is ready first.
- Makes concurrent job monitoring simple and readable.
*/
