package main

import (
   "fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	fmt.Printf("%s atoz[:9] = %s\n", atoz, atoz[:9])

	fmt.Printf("%s atoz[15:19] = %s\n", atoz, atoz[15:19])

    println("Ranging against atoz")

    for i, chr := range atoz {
    	fmt.Printf("%d=%c\n", i, chr)
    }

    fmt.Printf("len(atoz) %d\n", len(atoz))


    backquotes := `the "quick" brown fox jumps over the lazy dog\n`
    fmt.Printf("backquotes `%s`\n", backquotes)
}