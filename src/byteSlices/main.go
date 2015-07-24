package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("data/test.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	b := make([]byte, 100)

	n, err := f.Read(b)

	fmt.Printf("%d % x \n", n, b)

	stringVersion := string(b)

	fmt.Printf("\n%d - %s", n, stringVersion)

	someString := "foo bar"

	fout, err := os.Create("data/out.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer fout.Close()

	fout.Write([]byte(someString)) // here we convert/cast string to []byte before writing

}
