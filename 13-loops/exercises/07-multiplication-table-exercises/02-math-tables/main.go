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
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	v, err := strconv.Atoi(os.Args[2])
	o := os.Args[1]

	if err != nil {
		fmt.Println("Usage: [op=*/+-] [size]")
		return
	}

	if !strings.ContainsAny(o, "*/+-") {
		fmt.Println("Invalid operator. Valid ops one of: */+-")
	}

	fmt.Printf("%5s", o)

	for i := 0; i <= v; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()
	for i := 0; i <= v; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= v; j++ {
			switch o {
			case "*":
				fmt.Printf("%5d", i*j)
			case "/":
				if i == 0 || j == 0 {
					fmt.Printf("%5d", 0)
				} else {
					fmt.Printf("%5d", i/j)
				}

			case "+":
				fmt.Printf("%5d", i+j)
			case "-":
				fmt.Printf("%5d", i-j)
			}
		}
		fmt.Println()

	}

}

// AUTHOR SOLUTION
// const (
// 	validOps = "%*/+-"

// 	usageMsg       = "Usage: [op=" + validOps + "] [size]"
// 	sizeMissingMsg = "Size is missing"
// 	invalidOpMsg   = `Invalid operator.
// Valid ops one of: ` + validOps

// 	invalidOp = -1
// )

// func main() {
// 	args := os.Args[1:]

// 	switch l := len(args); {
// 	case l == 1:
// 		fmt.Println(sizeMissingMsg)
// 		fallthrough
// 	case l < 1:
// 		fmt.Println(usageMsg)
// 		return
// 	}

// 	op := args[0]
// 	if strings.IndexAny(op, validOps) == invalidOp {
// 		fmt.Println(invalidOpMsg)
// 		return
// 	}

// 	size, err := strconv.Atoi(args[1])
// 	if err != nil || size <= 0 {
// 		fmt.Println("Wrong size")
// 		return
// 	}

// 	fmt.Printf("%5s", op)
// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()

// 	for i := 0; i <= size; i++ {
// 		fmt.Printf("%5d", i)

// 		for j := 0; j <= size; j++ {
// 			var res int

// 			switch op {
// 			case "*":
// 				res = i * j
// 			case "%":
// 				if j != 0 {
// 					res = i % j
// 				}
// 			case "/":
// 				if j != 0 {
// 					res = i / j
// 				}
// 			case "+":
// 				res = i + j
// 			case "-":
// 				res = i - j
// 			}

// 			fmt.Printf("%5d", res)
// 		}
// 		fmt.Println()
// 	}
// }
