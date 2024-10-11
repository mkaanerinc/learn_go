package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))

	fmt.Println(futureValue)
}
