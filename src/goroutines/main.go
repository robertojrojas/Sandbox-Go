package main

import (
	"fmt"
	"math/rand"
)

func emit(wordChannel chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		wordChannel <- word
	}

	close(wordChannel)
}

func emitNoClose(wordChannel chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		wordChannel <- word
	}

}

func makeRandoms(intChannel chan int) {

	for {
		intChannel <- rand.Intn(1000)
	}

}

func makeID(intChannel chan int) {
	id := 0

	for {
		intChannel <- id
		id += 1
	}

}

func rangeOverWordChannel(wordChannel chan string) {

	for word := range wordChannel {
		fmt.Printf("%s ", word)
	}

	fmt.Println()
}

func receiveWord(wordChannel chan string) {
	word, ok := <-wordChannel

	if !ok { // ok will contain if the channel has been closed 'false' and word is 'nil'
		fmt.Printf("No more words - [%s]\n", word)
	} else {
		fmt.Printf("%s \n", word)
	}

}

func main() {

	wordChannel := make(chan string)

	go emit(wordChannel) //generates words

	//rangeOverWordChannel(wordChannel) // range over channel to receive words

	//Manually receive each word
	receiveWord(wordChannel) // once

	receiveWord(wordChannel) // again

	receiveWord(wordChannel) // again

	receiveWord(wordChannel) // again

	receiveWord(wordChannel) // this will not get anything b/c channel is closed and there nothing else to receive

	// we can also receive words from multiple emit goroutines
	//anotherWordChannel := make(chan string)

	// This will cause this error: The quick The brown fox quick brown fox fatal error: all goroutines are asleep - deadlock!
	// go emitNoClose(anotherWordChannel) //generates words
	// go emitNoClose(anotherWordChannel) //generates words

	// for anotherWord := range anotherWordChannel {
	// 	fmt.Printf("%s ", anotherWord)
	// }

	// randomsChan := make(chan int)

	// go makeRandoms(randomsChan)

	// for random := range randomsChan {
	// 	fmt.Printf("%d\n", random)
	// }

	idChan := make(chan int)
	go makeID(idChan)

	for range [4]int{} {
		fmt.Printf("ID: %d\n", <-idChan)
	}

}
