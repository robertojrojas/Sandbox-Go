package main

import (
	"fmt"
	"time"
)

func printer(msg string, goCh chan bool) {
	<-goCh
	fmt.Println("Starting to print...")

	fmt.Printf("%s\n", msg)
}

func stopPrinting(msg string, stopCh chan bool) {
	for {
		select {
			case <-stopCh:
				fmt.Println("Stopping....")
				return
			default:
				 fmt.Printf("%s ", msg)
		}
	}
}

func main() {

    goCh := make(chan bool)

	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer %d", i), goCh)
	}

	fmt.Println("Waiting 5 secs...")
	time.Sleep(5 * time.Second)

	close(goCh)

	time.Sleep(5 * time.Second)	

	stopCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go stopPrinting(fmt.Sprintf("."), stopCh)
	}

	time.Sleep(5 * time.Second)

	close(stopCh)

	time.Sleep(5 * time.Second)





}