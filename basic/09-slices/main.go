package main

import "fmt"

func main() {
	fmt.Println("Welcome to the lesson on Go Slices")

	// Declaration and Initialization
	// Slices are like dynamic arrays in Go
	// Declaration: var sliceName []Type
	// Initialization: sliceName := make([]Type, length, capacity)
	// Example:
	var mySlice []int           // Declaration
	mySlice = make([]int, 3, 5) // Initialization with length 3 and capacity 5
	fmt.Println("Empty Slice:", mySlice)

	// Initialization with literals
	mySlice = []int{1, 2, 3} // Initialization with literals
	fmt.Println("Initialized Slice:", mySlice)

	// Accessing elements
	fmt.Println("First Element:", mySlice[0])

	// Modifying elements
	mySlice[1] = 4
	fmt.Println("Modified Slice:", mySlice)

	// Appending elements
	mySlice = append(mySlice, 5, 6, 7)
	fmt.Println("After Appending:", mySlice)

	// Slicing existing slices
	// Syntax: sliceName[startIndex:endIndex]
	// startIndex is inclusive, endIndex is exclusive
	// Example:
	fmt.Println("Sliced Slice:", mySlice[1:3])

	// Length and Capacity
	fmt.Println("Length:", len(mySlice))
	fmt.Println("Capacity:", cap(mySlice))

	// Benefits of using slices compared to arrays:
	// - Slices are dynamic in size, allowing for flexible data structures.
	// - Slices are references to underlying arrays, so modifying a slice affects the original array.
	// - Slices are more memory-efficient than arrays, as they only allocate memory as needed.
	// - Slices support built-in operations like appending, slicing, and copying.

	// Comparison between array and slice
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]
	fmt.Println("Array:", array)
	fmt.Println("Slice:", slice)

	// Commonly used methods:
	// - append: Adds elements to the end of a slice
	// - copy: Copies elements from one slice to another
	// - len: Returns the length of the slice
	// - cap: Returns the capacity of the slice

	// Example of appending elements
	mySlice = append(mySlice, 8, 9)
	fmt.Println("After Appending Again:", mySlice)

	// Example of copying slices
	copiedSlice := make([]int, len(mySlice))
	copy(copiedSlice, mySlice)
	fmt.Println("Copied Slice:", copiedSlice)

	// Example 1: Slicing an existing array
	array1 := [5]int{1, 2, 3, 4, 5}
	slice1 := array1[1:4] // Slicing from index 1 to 3 (end index is exclusive)
	fmt.Println("Sliced Array:", slice1)

	// Example 2: Appending elements to a slice
	mySlice1 := []int{1, 2, 3}
	mySlice1 = append(mySlice1, 4, 5, 6) // Appending multiple elements
	fmt.Println("After Appending:", mySlice)

	// Example 3: Copying slices
	original := []int{1, 2, 3}
	copiedSlice1 := make([]int, len(original))
	copy(copiedSlice, original)
	fmt.Println("Copied Slice:", copiedSlice1)

	// Example 4: Creating a slice using the make function
	// Syntax: make([]Type, length, capacity)
	sliceWithMake := make([]int, 3, 5) // Length 3, Capacity 5
	fmt.Println("Slice with Make:", sliceWithMake)

	// Example 5: Length and Capacity of a slice
	anotherSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length of Slice:", len(anotherSlice))   // Length: 5
	fmt.Println("Capacity of Slice:", cap(anotherSlice)) // Capacity: 5

	// Example 6: Modifying a slice affects the original array
	arrayToSlice := [3]int{1, 2, 3}
	sliceFromArr := arrayToSlice[:] // Creating a slice from the entire array
	sliceFromArr[0] = 100           // Modifying the slice
	fmt.Println("Modified Slice:", sliceFromArr)
	fmt.Println("Original Array:", arrayToSlice)

	// Example 7: Iterating over a slice
	myNumbers := []int{10, 20, 30, 40, 50}
	for index, value := range myNumbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
