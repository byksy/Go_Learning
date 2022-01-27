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
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort the backing array
//
//  1. Sort only the middle 3 items.
//
//  2. All the slices should see your changes.
//
//
// RESTRICTION
//
//  Do not sort manually. Sort by using the sort package.
//
//
// EXPECTED OUTPUT
//
//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
//
// HINT:
//
//   Middle items are         : [frogger asteroids simcity]
//
//   After sorting they become: [asteroids frogger simcity]
//
// ---------------------------------------------------------

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	// ADD YOUR CODE HERE
	//MINE:
	// back := items[8:]
	// sorted := append([]string(nil), items[5:8]...)
	// sort.Strings(sorted)
	// items = append(items[0:5], sorted...)
	// items = append(items, back...)
	// AUTHOR SOLUTION
	mid := len(items) / 2
	smid := items[mid-1 : mid+2]

	// sorting the smid will affect the items
	// as well. their backing array is the same.
	sort.Strings(smid)
	fmt.Println()
	fmt.Println("Sorted  :", items)

}
