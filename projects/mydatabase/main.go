package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

// Define a Logger interface for logging purposes
type Logger interface {
	Fatal(string, ...interface{})
	Error(string, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Trace(string, ...interface{})
}

// Driver struct to manage the database operations
type Driver struct {
	mutex   sync.Mutex // Mutex for concurrency safety
	mutexes map[string]*sync.Mutex
	dir     string
	logger  Logger
}

// Options for configuring the Driver
type Options struct {
	Logger
}

// Address represents the user's address
type Address struct {
	City    string      `json:"city"`
	State   string      `json:"state"`
	Country string      `json:"country"`
	Pincode json.Number `json:"pincode"`
}

// User represents a user with personal and address information
type User struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Contact string  `json:"contact"`
	Company string  `json:"company"`
	Address Address `json:"address"`
}

// New creates a new instance of the Driver
func New(dir string, options *Options) (*Driver, error) {
	// Clean the directory path
	dir = filepath.Clean(dir)
	opts := Options{}

	// Set default options if not provided
	if options != nil {
		opts = *options
	}
	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}

	// Initialize the Driver instance
	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		logger:  opts.Logger,
	}

	// Check if the database directory exists
	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}
	opts.Logger.Debug("Creating the database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

// stat returns information about the file or directory at the given path
func stat(path string) (fi os.FileInfo, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

// Write writes data to a specified collection and resource
func (d *Driver) Write(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - unable to read!")
	}

	if resource == "" {
		return fmt.Errorf("missing resource - unable to read (no name)!")
	}

	// Lock the mutex for the collection to prevent concurrent writes
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	// Create directories if they don't exist
	dir := filepath.Join(d.dir, collection)
	finalPath := filepath.Join(dir, resource+".json")
	tempPath := finalPath + ".temp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Marshal data to JSON format
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	b = append(b, byte('\n'))

	// Write data to a temporary file first
	if err := ioutil.WriteFile(tempPath, b, 0644); err != nil {
		return err
	}

	// Rename the temporary file to the final path
	return os.Rename(tempPath, finalPath)
}

// Read reads data from a specified collection and resource
func (d *Driver) Read(collection, resource string, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("missing collection - unable to read!")
	}
	if resource == "" {
		return fmt.Errorf("missing resource - unable to read (no name)!")
	}

	// Construct the file path
	record := filepath.Join(d.dir, collection, resource)

	// Check if the file exists
	if _, err := stat(record); err != nil {
		return err
	}

	// Read data from the file
	b, err := ioutil.ReadFile(record + ".json")

	if err != nil {
		return err
	}

	// Unmarshal JSON data into the provided interface
	return json.Unmarshal(b, &v)
}

// ReadAll reads all records from a specified collection
func (d *Driver) ReadAll(collection string) ([]string, error) {
	if collection == "" {
		return nil, fmt.Errorf("missing collection - unable to read!")
	}

	// Construct the directory path
	dir := filepath.Join(d.dir, collection)

	// Check if the directory exists
	if _, err := stat(dir); err != nil {
		return nil, err
	}

	// Read files from the directory
	files, _ := ioutil.ReadDir(dir)

	var records []string

	// Read data from each file and append to the records slice
	for _, file := range files {
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))

		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}

	return records, nil
}

// Delete deletes a specified resource from a collection
func (d *Driver) Delete(collection, resource string) error {
	path := filepath.Join(collection, resource)
	if _, err := stat(path); err != nil {
		return err
	}
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, path)

	switch fi, err := stat(dir); {
	case fi == nil, err != nil:
		return fmt.Errorf("unable to find file or directory named %v", dir)
	case fi.Mode().IsDir():
		return os.RemoveAll(dir)
	case fi.Mode().IsRegular():
		return os.RemoveAll(dir + ".json")
	}
	return nil
}

// getOrCreateMutex returns a mutex for the given collection
func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}

func main() {
	// Set the directory for the database
	dir := "./collections"

	// Initialize the database
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// Initialize example data for employees
	employees := []User{
		{Name: "John", Age: 23, Contact: "9832098320", Company: "Google", Address: Address{City: "Bangalore", State: "Karnataka", Country: "India", Pincode: "410013"}},
		{Name: "Alice", Age: 25, Contact: "9876543210", Company: "Microsoft", Address: Address{City: "Mumbai", State: "Maharashtra", Country: "India", Pincode: "400001"}},
		{Name: "Bob", Age: 28, Contact: "9765432109", Company: "Amazon", Address: Address{City: "Delhi", State: "Delhi", Country: "India", Pincode: "110001"}},
		{Name: "Emma", Age: 30, Contact: "9898765432", Company: "Facebook", Address: Address{City: "Hyderabad", State: "Telangana", Country: "India", Pincode: "500001"}},
		{Name: "David", Age: 35, Contact: "9876543210", Company: "Apple", Address: Address{City: "Pune", State: "Maharashtra", Country: "India", Pincode: "411001"}},
	}

	// Write employee data to the database
	for _, emp := range employees {
		db.Write("users", emp.Name, User{
			Name:    emp.Name,
			Age:     emp.Age,
			Contact: emp.Contact,
			Company: emp.Company,
			Address: emp.Address,
		})
	}

	// Read all records from the "users" collection
	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// Print all records
	fmt.Println(records)

	// Unmarshal records into User struct
	allUsers := []User{}
	for _, f := range records {
		employFound := User{}
		if err := json.Unmarshal([]byte(f), &employFound); err != nil {
			fmt.Println("Error", err)
			return
		}
		allUsers = append(allUsers, employFound)
	}

	// Print all users
	fmt.Println(allUsers)

	// Delete a user record
	// if err := db.Delete("users", "John"); err != nil {
	// 	fmt.Println("Error", err)
	// }
}
