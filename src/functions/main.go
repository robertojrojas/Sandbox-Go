package main

import (
	"fmt"
	"os"
	"time"
)

func printer(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf(msg)
	return msg, err
}

func printerDefer(msg string) error {
	defer fmt.Printf("--\n")
	defer fmt.Printf("More\n")

	_, err := fmt.Printf("%s", msg)

	return err
}

func filePrinter(msg ...string) (e error) {
	f, e := os.Create("Helloworld.txt")
	if e != nil {
		return
	}
    defer f.Close()


    for _, msgToWrite := range msg {

    	msgToWrite := fmt.Sprintf("%s - %s\n", time.Now().Local(), msgToWrite)
    	f.Write([]byte(msgToWrite))	

    }

    return

}

func main() {

	var original = "Hello, World!"

	appendedMessage, err := printer(original)
	if err == nil {
		fmt.Printf("%q\n", appendedMessage)
	}
	fmt.Printf("Original %s\n", original)

	printerDefer(original)

	filePrinter(original, "Hola Roberto!")
	// filePrinter() // this will create an empty file since no strings
}