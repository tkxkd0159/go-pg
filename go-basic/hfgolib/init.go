package hfgolib

import "fmt"

// init is initial function to be executed prior to the start hfgolib
// For multiple init in a single file, their processing order is in the order of their declaration,
// while init declared in multiple files are processed according to the lexicographic filename order.
func init() {
	fmt.Println("I am hfgolib")
}
