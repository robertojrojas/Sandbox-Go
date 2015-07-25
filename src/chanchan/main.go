package main

import (
	"fmt"
)

var (
	words = []string {"The", "quick", "brown", "fox", "jumps"}
)

func emit(wordChanChan chan chan string, done chan bool) {

	wordChannel := make(chan string)

	wordChanChan <- wordChannel
	
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


func main() {

    wordChanChan := make(chan chan string)
	doneCh := make(chan bool)

    go emit(wordChanChan, doneCh)

    wordCh := <- wordChanChan
	

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



}