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
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------
const (
	max_rand     = 5
	user_message = "Wrong usage"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println(user_message)
		return
	}

	guess1, err1 := strconv.Atoi(args[0])
	guess2, err2 := strconv.Atoi(args[1])

	if err1 != nil || err2 != nil {
		fmt.Println(user_message)
		return
	}

	for turn := 1; turn < max_rand; turn++ {

		n := rand.Intn(guess2 + 1)

		if n == guess1 || n == guess2 {
			fmt.Println("You won!!!")
			return
		}
	}

	fmt.Println("You lost!!!")
}

// AUTHOR SOLUTION

// const (
// 	maxTurns = 5 // less is more difficult
// 	usage    = `Welcome to the Lucky Number Game! 🍀
// The program will pick %d random numbers.
// Your mission is to guess one of those numbers.
// The greater your number is, harder it gets.
// Wanna play?
// `
// )

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	args := os.Args[1:]

// 	if len(args) < 1 {
// 		fmt.Printf(usage, maxTurns)
// 		return
// 	}

// 	guess, err := strconv.Atoi(args[0])
// 	if err != nil {
// 		fmt.Println("Not a number.")
// 		return
// 	}

// 	var guess2 int
// 	if len(args) == 2 {
// 		guess2, err = strconv.Atoi(args[1])
// 		if err != nil {
// 			fmt.Println("Not a number.")
// 			return
// 		}
// 	}

// 	if guess <= 0 || guess2 <= 0 {
// 		fmt.Println("Please pick positive numbers.")
// 		return
// 	}

// 	min := guess
// 	if guess < guess2 {
// 		min = guess2
// 	}

// 	for turn := 0; turn < maxTurns; turn++ {
// 		n := rand.Intn(min) + 1

// 		if n == guess || n == guess2 {
// 			fmt.Println("🎉  YOU WIN!")
// 			return
// 		}
// 	}

// 	fmt.Println("☠️  YOU LOST... Try again?")
// }
