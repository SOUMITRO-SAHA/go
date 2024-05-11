package main

import "fmt"

func main() {
	// Welcome message
	fmt.Println("Welcome to the Go Arrays lesson!")

	// Declaring an array with a fixed size of 4
	var firstArray [4]string

	// Inserting elements into the array
	firstArray[0] = "zero"
	firstArray[1] = "one"

	// Printing the array and its length
	fmt.Println("First Array:", firstArray)
	fmt.Println("Length of the First Array:", len(firstArray))

	// Declaring and initializing an array in one line
	var secondArray = [4]string{"two", "three", "four", "five"}

	// Printing the second array and its length
	fmt.Println("Secondary Array:", secondArray)
	fmt.Println("Length of the Secondary Array:", len(secondArray))

	// Accessing individual elements of the array
	fmt.Println("Element at index 2 of the second array:", secondArray[2])

	// Iterating over the elements of the array using a for loop
	fmt.Println("Iterating over the elements of the second array:")
	for i := 0; i < len(secondArray); i++ {
		fmt.Printf("Index %d: %s\n", i, secondArray[i])
	}

	// Modifying an element of the array
	secondArray[2] = "four_modified"
	fmt.Println("Modified second array:", secondArray)

	// Initializing an array with default values
	var thirdArray [5]int // All elements will be initialized to their zero value (0 for int)
	fmt.Println("Third array (initialized with default values):", thirdArray)

	// Initializing an array with specific values for some elements
	fourthArray := [5]int{1, 2} // Only the first two elements will be 1 and 2, the rest will be initialized to 0
	fmt.Println("Fourth array (initialized with specific values):", fourthArray)
}
