package main

import (
	"fmt"
	"time"
)

var (
	words = []string {"The", "quick", "brown", "fox", "jumps"}
)

func emit(wordChannel chan string, done chan bool) {
	
    i := 0
   
	for {
		fmt.Printf("\nfor...\n")
		select {
		case wordChannel <- words[i]:
			i += 1
			if i == len(words) {
				i = 0
			}
		case <-done:
			fmt.Printf("\nGot done!\n")
			done <- true
			return
		}
	}
}

func emitTimeout(wordChannel chan string, done chan bool) {
	//defer close(wordChannel)

    i := 0

    t := time.NewTimer(3 * time.Second)
   
	for {
	
		select {
			case wordChannel <- words[i]:
				i += 1
				if i == len(words) {
					i = 0
				}
			case <-done:
				fmt.Printf("\nGot done!\n")
				done <- true
				return

			case <-t.C:
				close(wordChannel) //or we can use the 'defer close(wordChannel)' above
				return
		}
	}
}

func main() {

	wordCh := make(chan string)
	doneCh := make(chan bool)

    go emit(wordCh, doneCh)

    modValue := 0
    modBreak := len(words)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
		modValue++

		if i == 0 {
			continue
		}

		if (modValue == modBreak) {
			fmt.Println()
			modValue = 0
		}
	}

    // This is used as bi-directional done, confirmed communication technique between goroutines
	doneCh <- true // Tell it to stop
	<- doneCh // Confirmation that it stopped



    wordTimeoutCh := make(chan string)
	doneTimeoutCh := make(chan bool)

    go emitTimeout(wordTimeoutCh, doneTimeoutCh)

    for word := range wordTimeoutCh {
    	fmt.Printf("%s \n", word)
    }



}