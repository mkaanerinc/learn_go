package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/price-calculator/conversion"
)

// TaxIncludedPriceJob represents a job for calculating tax-included prices.
type TaxIncludedPriceJob struct {
	TaxRate           float64            // Tax rate to apply (e.g., 0.18 for 18%)
	InputPrices       []float64          // Raw input prices before tax
	TaxIncludedPrices map[string]float64 // Optional: stores tax-included prices (currently unused)
}

// LoadData reads price data from a text file and populates the InputPrices slice.
func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt") // Attempt to open the file containing raw prices

	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file) // Scanner to read file line by line
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Collect each line into the lines slice
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed")
		fmt.Println(err)
		file.Close()
		return
	}

	// Convert each line from string to float64; fail fast on invalid input
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println("Converting price to float failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	job.InputPrices = prices // Store parsed prices in the job struct
	file.Close()
}

// Process loads input prices and calculates tax-included prices, printing them as a map.
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData() // Load prices from file

	result := make(map[string]string) // Create a map to hold formatted price mappings

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)                              // Apply tax rate to each price
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice) // Format and store
	}

	fmt.Println(result)
}

// NewTaxIncludedPriceJob is a constructor function that initializes a new job instance.
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30}, // Default hardcoded prices (used only if LoadData isn't called)
	}
}
