package main

import (
	"fmt"
)

// Arrays are passed by value
func printer(w [9]string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	w[2] = "Blue" // This change will not affect the caller
}

// Slices are passed by reference
func printerSlice(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Printf("\n")

	w[2] = "Blue" // This change will affect the caller
}

func main() {

	words := [...]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	fmt.Printf("Printing arrays\n")
	printer(words)
	printer(words)

	wordsSlice := words[:] // convert array to slice

	fmt.Printf("Printing arrays\n")
	printerSlice(wordsSlice)
	printerSlice(wordsSlice)

	printerSlice(words[5:len(words)])
	printerSlice(words[5:])
	printerSlice(words[:3])

	otherWords := make([]string, 4)
	otherWords[0] = "The"
	otherWords[1] = "Quick"
	otherWords[2] = "Brown"
	otherWords[3] = "Fox"
	printerSlice(otherWords)

	moreWords := make([]string, 0, 4)
	fmt.Printf("starting - len(moreWords) = %d  cap(moreWords) = %d\n", len(moreWords), cap(moreWords))
	moreWords = append(moreWords, "The")
	moreWords = append(moreWords, "Quick")
	moreWords = append(moreWords, "Brown")
	moreWords = append(moreWords, "Fox")
	fmt.Printf("Addded first items - len(moreWords) = %d  cap(moreWords) = %d\n", len(moreWords), cap(moreWords))
	printerSlice(moreWords)
	moreWords = append(moreWords, "Jumps")
	fmt.Printf("Added one more item - len(moreWords) = %d  cap(moreWords) = %d\n", len(moreWords), cap(moreWords))
	printerSlice(moreWords)

	newWords := make([]string, len(moreWords)) // the new slice needs to be at least the same len as the source slice
	copy(newWords, moreWords)                  // since the second argument is a slice we could have used moreWords[:2] to just copy the first to items
	newWords[2] = "Red"
	fmt.Printf("After copy and modification\n")
	printerSlice(moreWords)
	printerSlice(newWords)
	fmt.Printf("len(newWords) = %d  cap(newWords) = %d\n", len(newWords), cap(newWords))

}
