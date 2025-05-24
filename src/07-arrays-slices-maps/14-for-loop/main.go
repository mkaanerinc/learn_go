package main

import "fmt"

type float64Map map[string]float64

func (m float64Map) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Kaan")
	userNames = append(userNames, "Mustafa")

	fmt.Println(userNames)

	courseRating := make(float64Map, 3)
	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["angular"] = 4.8

	courseRating.output()

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRating {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
