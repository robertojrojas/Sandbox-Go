package main

import (
	"errors"
	"fmt"
)

var (
	errorEmptyString = errors.New("CustomError - Unwilling to print an empty stirng")
)

func printer(msg string) error {

	if msg == "" {
		return fmt.Errorf("Unwilling to print an empty stirng")
	}
	_, err := fmt.Printf("%s\n", msg)

	return err
}

func printerCustomError(msg string) error {

	if msg == "" {
		return errorEmptyString
	}
	_, err := fmt.Printf("%s\n", msg)

	return err
}

func printerPanic(msg string) error {

	if msg == "" {
		panic(errorEmptyString)
	}
	_, err := fmt.Printf("%s\n", msg)

	return err
}

func main() {

	if err := printer(""); err != nil {
		fmt.Printf("Printer failed: %s\n", err)
		//os.Exit(1)
	}

	if err := printerCustomError(""); err != nil {
		if err == errorEmptyString {
			fmt.Printf("You tried to print an empty string!\n")
		} else {
			fmt.Printf("Printer failed: %s\n", err)
		}

		//os.Exit(1)
	}

	if err := printerPanic(""); err != nil {
		fmt.Printf("Printer failed: %s\n", err)
		//os.Exit(1)
	}

}
