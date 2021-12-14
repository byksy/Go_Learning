// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func main() {
	// BONUS: Use a variable for the format specifier
	var (
		firstName = "Mustafa"
		lastName  = "BÜYÜKSOY"
	)

	fmt.Printf("My name is %s and my last name is %s. \n", firstName, lastName)

	msg := "My name is %s and my last name is %s. \n"
	fmt.Printf(msg, "Mustafa", "BÜYÜKSOY")

	// fmt.Printf("?", ?, ?)
}
