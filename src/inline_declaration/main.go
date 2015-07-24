package main

import (
	"fmt"
	"os"
)

func main() {

	// numberOfBytes and err are scoped within the 'if'
	if numberOfBytes, err := fmt.Printf("Hello, World!\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d characters\n", numberOfBytes)
	}

	var numberOfBytes int //this is a separate declaration than then one above
	var err error
	if numberOfBytes, err = fmt.Printf("Hello, World!\n"); err != nil {
		os.Exit(1)
	}

    fmt.Printf("Outside declariation of nubmerOfBytes, Printed %d characters\n", numberOfBytes)
	
}