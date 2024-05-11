package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Lesson of Go Map (also known as HashMap in Java)")

	// Example 1: Declaring and initializing a map
	var myMap map[string]int // Declaring a map (key: string, value: int)
	myMap = make(map[string]int)
	myMap["apple"] = 10
	myMap["banana"] = 20
	myMap["cherry"] = 30
	fmt.Println("Initialized Map:", myMap)

	// Example 2: Accessing elements in a map
	appleCount := myMap["apple"]
	fmt.Println("Count of Apples:", appleCount)

	// Example 3: Modifying elements in a map
	myMap["banana"] = 25 // Modifying the value for the key "banana"
	fmt.Println("Modified Map:", myMap)

	// Example 4: Deleting elements from a map
	delete(myMap, "cherry") // Deleting the key "cherry" from the map
	fmt.Println("Map after Deletion:", myMap)

	// Example 5: Checking if a key exists in the map
	_, isPresent := myMap["banana"]
	fmt.Println("Is 'banana' present in the map?", isPresent)

	// Example 6: Iterating over a map
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// Example 7: Length of a map
	fmt.Println("Length of Map:", len(myMap))

	// Example 8: Creating a map with initial values
	initialMap := map[string]string{
		"one":   "first",
		"two":   "second",
		"three": "third",
	}
	fmt.Println("Map with Initial Values:", initialMap)
}
