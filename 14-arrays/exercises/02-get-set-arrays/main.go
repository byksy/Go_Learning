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
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
// ---------------------------------------------------------

func main() {

	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)

	names[0] = "Mustafa"
	names[1] = "Asya"
	names[2] = "Nevra"
	separator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("names", separator)
	for i := range names {

		fmt.Printf("names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\ndistances", separator)

	distances[0] = 50
	distances[1] = 50
	distances[2] = 50
	distances[3] = 50
	distances[4] = 50

	for i := range distances {

		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\ndata", separator)

	data[0] = 1
	data[1] = 2
	data[2] = 3
	data[3] = 4
	data[4] = 5

	for i := range data {

		fmt.Printf("data[%d]: %d\n", i, data[i])
	}

	for i, v := range data {
		// try the %c verb
		fmt.Printf("data[%d]: %c\n", i, v)
	}
	fmt.Print("\nratios", separator)

	ratios[0] = 3.14

	fmt.Printf("ratios[0]: %.2f\n", ratios[0])

	fmt.Print("\nalives", separator)

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	for i := range alives {

		fmt.Printf("alives[%d], %t\n", i, alives[i])
	}

	fmt.Print("\nzero", separator)

	fmt.Printf("zero: %d\n", zero)
}
