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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// Solve this exercise by using your previous solution for
	// the "Housing Prices" exercise.

	head := strings.Split(header, ",")

	locations, size, beds, baths, price := []string{}, []int{}, []int{}, []int{}, []int{}

	rows := strings.Split(data, "\n")

	for _, row := range rows {

		cols := strings.Split(row, separator)

		locations = append(locations, cols[0])

		n, _ := strconv.Atoi(cols[1])
		size = append(size, n)

		n, _ = strconv.Atoi(cols[2])
		beds = append(beds, n)

		n, _ = strconv.Atoi(cols[3])
		baths = append(baths, n)

		n, _ = strconv.Atoi(cols[4])
		price = append(price, n)

	}

	for h := range head {
		fmt.Printf("%-15s", head[h])
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 70))

	var (
		sum_size, sum_beds, sum_baths, sum_price float64
	)
	for h := range rows {
		fmt.Printf("%-15s", locations[h])
		fmt.Printf("%-15d", size[h])
		fmt.Printf("%-15d", beds[h])
		fmt.Printf("%-15d", baths[h])
		fmt.Printf("%-15d", price[h])
		sum_size += float64(size[h])
		sum_beds += float64(beds[h])
		sum_baths += float64(baths[h])
		sum_price += float64(price[h])
		fmt.Println()
	}

	fmt.Printf("\n%s\n", strings.Repeat("=", 70))
	fmt.Printf("%-15s", " ")
	fmt.Printf("%-15.2f", sum_size/4)
	fmt.Printf("%-15.2f", sum_beds/4)
	fmt.Printf("%-15.2f", sum_baths/4)
	fmt.Printf("%-15.2f", sum_price/4)

}
