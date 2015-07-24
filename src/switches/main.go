package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!")

	switch {
		case err != nil:
			os.Exit(1)
		case n == 0:
			fmt.Printf("No bytes output")
		case n != 13:
			fmt.Printf("Wrong number of characters: %d", n)
		default:
			fmt.Printf("OK!")
	}

	fmt.Printf("\n")

    atoz := "the quick brown fox jumps over the lazy dog"
    vowels := 0
    consonants := 0
    zeds := 0

	for _, char := range atoz {
		switch char {
			case 'a', 'e', 'i', 'o', 'u':
				vowels += 1
			case 'z':
				zeds += 1
				fallthrough //since 'z' is a consonant
			default:
				consonants += 1

		}
	}

	fmt.Printf("vowels %d consonants %d zeds %d\n", vowels, consonants, zeds)


}