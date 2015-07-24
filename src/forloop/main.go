package main

import (
	"fmt"
)

func main() {

	// var counter int 

	// counter = 0
	// for counter < 10 {

    for i, j := 0, 1; i < 10; i, j = i+1, j*2 {
      	fmt.Printf("%d Hello, World!\n", j)
	}
}