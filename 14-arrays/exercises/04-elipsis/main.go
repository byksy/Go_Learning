package main

import (
	"fmt"
	"strings"
)

func main() {

	names := [...]string{"Mustafa", "Asya", "Nevra"}

	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)

	for i := range names {

		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)

	distances := [...]int{50, 50, 30, 20, 10}

	for i := range distances {

		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)

	data := [...]uint8{1, 2, 3, 4, 5}

	for i := range data {

		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %c\n", i, v)
	}
	fmt.Print("\nratios", separator)

	ratios := [...]float64{3.14}

	fmt.Printf("ratios[0]: %.2f\n", ratios[0])

	fmt.Print("\nalives", separator)

	alives := [...]bool{true, false, true, false}

	for i := range alives {

		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}

	fmt.Print("\nzero", separator)
	// A byte array that doesn't occupy memory space
	//
	// Don't do this:
	//   zero := [0]byte{}
	//
	// Do this (when you don't assign elements):
	var zero [0]byte
	fmt.Printf("zero: %d\n", zero)
}
