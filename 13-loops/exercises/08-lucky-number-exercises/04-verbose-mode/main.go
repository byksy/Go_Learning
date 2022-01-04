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
// EXERCISE: Verbose Mode
//
//  When the player runs the game like this:
//
//    go run main.go -v 4
//
//  Display each generated random number:

//    1 3 4 🎉  YOU WIN!
//
//  In this example, computer picks 1, 3, and 4. And the
//  player wins.
//
// HINT
//  You need to get and interpret the command-line arguments.
// ---------------------------------------------------------

const (
	max_turn     = 10
	user_message = `Welcome to the Lucky Number Game! 🍀
	The program will pick %d random numbers.
	Your mission is to guess one of those numbers.
	The greater your number is, harder it gets.
	Wanna play?`
)

func main() {

	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Printf(user_message, max_turn)
		return
	}

	v := args[0]
	guess, err := strconv.Atoi(args[1])

	if err != nil {
		fmt.Printf(user_message, max_turn)
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 1; turn < max_turn; turn++ {
		n := rand.Intn(guess + 1)

		if v == "-v" {
			fmt.Printf("%d ", n)
			if n == guess {
				fmt.Println("🎉 You win")
				return
			}
		}
	}

	fmt.Println("😢 You lost!")
}
