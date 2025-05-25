package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
)

// TaxIncludedPriceJob represents a job for calculating tax-included prices.
type TaxIncludedPriceJob struct {
	IOManager         filemanager.FileManager
	TaxRate           float64           // Tax rate to apply (e.g., 0.18 for 18%)
	InputPrices       []float64         // Raw input prices before tax
	TaxIncludedPrices map[string]string // Optional: stores tax-included prices (currently unused)
}

// LoadData reads price data from a text file and populates the InputPrices slice.
func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert each line from string to float64; fail fast on invalid input
	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices // Store parsed prices in the job struct
}

// Process loads input prices and calculates tax-included prices, printing them as a map.
func (job *TaxIncludedPriceJob) Process() {
	job.LoadData() // Load prices from file

	result := make(map[string]string) // Create a map to hold formatted price mappings

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)                              // Apply tax rate to each price
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice) // Format and store
	}

	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
}

// NewTaxIncludedPriceJob is a constructor function that initializes a new job instance.
func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   fm,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30}, // Default hardcoded prices (used only if LoadData isn't called)
	}
}
