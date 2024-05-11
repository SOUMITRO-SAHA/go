package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Go Pointers Lesson")

	// Declare a variable 'value' and initialize it with 10
	var value int = 10
	fmt.Println("Value of the Variable:", value)

	// Print the memory address of the variable 'value'
	fmt.Println("Address of the Variable:", &value)

	// Declare a pointer variable 'ptr' of type integer
	var ptr *int

	// Assign the memory address of the variable 'value' to the pointer 'ptr'
	ptr = &value

	// Print the value stored at the memory address pointed to by 'ptr'
	fmt.Println("Value of the Pointer (De-Referencing):", *ptr)

	// Print the memory address stored in the pointer 'ptr'
	fmt.Println("Address stored in the Pointer:", ptr)

	// Increment the value stored at the memory address pointed to by 'ptr' by 10
	*ptr = *ptr + 10

	// The value of 'value' is modified because 'ptr' points to the same memory address as 'value'
	fmt.Println("Value of the Variable 'value' after modification:", value)

	// ########## [ Passing Pointers to a Function in Go ] ######### //
	fmt.Println("Passing Pointers to a Function in Go")

	// Declare a variable 'num' and initialize it with 10
	num := 10
	fmt.Println("Value before increment:", num)

	// Pass the memory address of 'num' to the function 'incrementByPointer'
	incrementByPointer(&num)

	// Print the updated value of 'num'
	fmt.Println("Value after increment:", num)

	// ########## [ Pointer to a Struct in Golang ] ######### //

	fmt.Println("Pointer to a Struct in Go")

	// Create a new instance of the Person struct
	person := Person{Name: "Alice", Age: 30}

	// Create a pointer to the Person struct
	ptrToPerson := &person

	// Print the original values
	fmt.Println("Original Person:", person)

	// Modify the Person struct using the pointer
	modifyPerson(ptrToPerson, "Bob", 35)

	// Print the modified values
	fmt.Println("Modified Person:", person)

}

// Function to increment the value of a variable using pointers
func incrementByPointer(ptr *int) {
	// Increment the value stored at the memory address pointed to by 'ptr'
	*ptr = *ptr + 1
}

// Define a struct type
type Person struct {
	Name string
	Age  int
}

// Function to modify a Person struct using a pointer
func modifyPerson(p *Person, name string, age int) {
	p.Name = name
	p.Age = age
}
