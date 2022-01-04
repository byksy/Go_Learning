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
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------
const (
	max_turns = 6
	usage     = `Welcome to the Lucky Number Game! 🍀
	The program will pick %d random numbers.
	Your mission is to guess one of those numbers.
	The greater your number is, harder it gets.
	Wanna play?`
)

func main() {

	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Printf(usage, max_turns)
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Printf(usage, max_turns)
		return
	}

	if guess <= 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for i := 1; i < max_turns; i++ {

		n := rand.Intn(guess + 1)

		if guess == n {
			if i == 1 {
				fmt.Println("Your are the first turn winner!!!")
			} else {
				fmt.Println("You win!!!")
			}

			return
		}

		// YOU CAN ALSO USE THE BELOW USAGE FOR BETTER SOLUTION

		// n := rand.Intn(guess) + 1

		// Better, why?
		//
		// Instead of nesting the if statement into
		//   another if statement; it simply continues.
		//
		// TLDR: Avoid nested statements.
		// if n != guess {
		// 	continue
		// }

		// if turn == 1 {
		// 	fmt.Println("🥇 FIRST TIME WINNER!!!")
		// } else {
		// 	fmt.Println("🎉  YOU WON!")
		// }
		// return

	}
	fmt.Println("You lost!!!")

}
