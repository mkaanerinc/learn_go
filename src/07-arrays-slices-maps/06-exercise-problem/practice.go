// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	// 1) Create a new array with three hobbies and output the array
	hobbies := [3]string{"Reading", "Cycling", "Cooking"}
	fmt.Println("Hobbies Array:", hobbies)

	// 2) Output First element and second-third element combined as a new list
	fmt.Println("First Hobby:", hobbies[0])
	fmt.Println("Second and Third Hobbies Combined:", hobbies[1:3])

	// 3) Create a slice based on the first element that contains the first and second elements.
	// reate that slice in two different ways (i.e. create two slices in the end)
	hobbiesSlice := hobbies[0:2]
	hobbiesSliceAlternative := hobbies[:2]
	fmt.Println("Hobbies Slice (First and Second):", hobbiesSlice)
	fmt.Println("Hobbies Slice Alternative:", hobbiesSliceAlternative)

	// 4) Re-slice the slice from (3) and change it to contain the second and last element of the original array.
	fmt.Println("Capacity of Hobbies Slice:", cap(hobbiesSlice))
	ReSliceHobbiesSlice := hobbiesSlice[1:3]
	fmt.Println("ReSliceHobbiesSlice:", ReSliceHobbiesSlice)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	courseGoals := []string{"Learn Go", "Build a Web Application"}

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	courseGoals[1] = "Build a REST API"
	courseGoals = append(courseGoals, "Contribute to Open Source")
	fmt.Println("Course Goals:", courseGoals)

	// 7) Create a "Product" struct with title, id, price create a dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	products := []Product{
		{title: "Go Programming Book", id: 1, price: 29.99},
		{title: "Web Development Course", id: 2, price: 49.99},
	}

	products = append(products, Product{title: "API Design Guide", id: 3, price: 19.99})
	fmt.Println("Products:", products)
}
