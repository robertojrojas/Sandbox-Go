package main

import "fmt"

func main() {

	// Dynamic map
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])

	// This returns 0
	fmt.Printf("Days in Wrong month: %d\n", dayMonths["wrong"])

	wrongMonth := "WrongMonth"
	days, ok := dayMonths[wrongMonth] // Since "wrongMonth" is not found ok will be false
	if !ok {
		fmt.Printf("Can't get days for '%s'\n", wrongMonth)
	} else {
		fmt.Printf("Days in correct: %d\n", days)
	}

	for month, days := range dayMonths {
		fmt.Printf("%s has %d days \n", month, days)
	}

	has31Days := 0
	has30Days := 0
	hasLessThan30Days := 0

	for _, days := range dayMonths {
		switch days {
		case 31:
			has31Days++
		case 30:
			has30Days++
		default:
			hasLessThan30Days++

		}
	}
	fmt.Printf("%d months have 31 days, %d months have 30 days, %d months have less than 30 days\n", has31Days, has30Days, hasLessThan30Days)

	delete(dayMonths, "Feb")
	delete(dayMonths, "Feb") // Can call delete twice!

	fmt.Println("After deleting Feb")
	_, isFebPresent := dayMonths["Feb"]
	if !isFebPresent {
		fmt.Printf("Feb was deleted\n")
	}

	// Static map
	otherDayMonths := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	fmt.Printf("Jan has %d days\n", otherDayMonths["Jan"])

}
