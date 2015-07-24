package main

import (
	"fmt"
)

func main() {
	var pi float64 = 3.14
    fmt.Printf("Float Value: %.2f\n", pi)

    nine := uint8(9)
    fmt.Printf("Int Value is %d\n", nine)

    var isTrue bool
    //isTrue = true

    fmt.Printf("Bool Value: %t\n", isTrue)


    b := byte(65)
    fmt.Printf("Byte value %v Hex value: %x\n", b, b)

}