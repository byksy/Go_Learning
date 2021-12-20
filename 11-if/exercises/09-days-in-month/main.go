// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Print the number of days in a given month.
//
// RESTRICTIONS
//  1. On a leap year, february should print 29. Otherwise, 28.
//
//     Set your computer clock to 2020 to see whether it works.
//
//  2. It should work case-insensitive. See below.
//
//     Search on Google: golang pkg strings ToLower
//
//  3. Get the current year using the time.Now()
//
//     Search on Google: golang pkg time now year
//
//
// EXPECTED OUTPUT
//
//  -----------------------------------------
//  Your solution should not accept invalid months
//  -----------------------------------------
//  go run main.go
//    Give me a month name
//
//  go run main.go sheep
//    "sheep" is not a month.
//
//  -----------------------------------------
//  Your solution should handle the leap years
//  -----------------------------------------
//  go run main.go january
//    "january" has 31 days.
//
//  go run main.go february
//    "february" has 28 days.
//
//  go run main.go march
//    "march" has 31 days.
//
//  go run main.go april
//    "april" has 30 days.
//
//  go run main.go may
//    "may" has 31 days.
//
//  go run main.go june
//    "june" has 30 days.
//
//  go run main.go july
//    "july" has 31 days.
//
//  go run main.go august
//    "august" has 31 days.
//
//  go run main.go september
//    "september" has 30 days.
//
//  go run main.go october
//    "october" has 31 days.
//
//  go run main.go november
//    "november" has 30 days.
//
//  go run main.go december
//    "december" has 31 days.
//
//  -----------------------------------------
//  Your solution should be case insensitive
//  -----------------------------------------
//  go run main.go DECEMBER
//    "DECEMBER" has 31 days.
//
//  go run main.go dEcEmBeR
//    "dEcEmBeR" has 31 days.
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 2 {

		fmt.Println("Give me a month name")
		return
	}
	m := strings.ToUpper(os.Args[1])

	// var leap = false
	// if time.Now().Year()%4 == 0 && (time.Now().Year()%400 == 0 || time.Now().Year()%100 != 0) {
	// 	leap = true
	// } else {
	// 	leap = false
	// }

	//Simple statement for leap

	leap := time.Now().Year()%4 == 0 && (time.Now().Year()%400 == 0 || time.Now().Year()%100 != 0)

	if m == "JANUARY" ||
		m == "MARCH" ||
		m == "MAY" ||
		m == "JULY" ||
		m == "AUGUST" ||
		m == "OCTOBER" ||
		m == "DECEMBER" {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if m == "FEBRUARY" && leap {
		fmt.Printf("%q has 29 days.\n", os.Args[1])

	} else if m == "FEBRUARY" && !leap {
		fmt.Printf("%q has 28 days.\n", os.Args[1])
	} else if m == "APRIL" ||
		m == "JUNE" ||
		m == "SEPTEMBER" ||
		m == "NOVEMBER" {
		fmt.Printf("%q has 30 days.\n", os.Args[1])
	} else {
		fmt.Printf("%q is not a month.\n", os.Args[1])
	}

}
