package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {

// MY SOLUTION

// feet := os.Args[1]

// n, err := strconv.Atoi(feet)

// if err != nil {

// 	fmt.Printf("error: %q is not a number.\n", feet)
// 	return
// }

// meter := float64(n) * 0.3048

// fmt.Printf("%d, feet is %.2f meters", n, meter)

// AUTHOR SOLUTION

// arg := os.Args[1]

// feet, err := strconv.ParseFloat(arg, 64)
// if err != nil {
// 	fmt.Printf("error: %q is not a number. \n", arg)
// 	return
// }
// meters := feet * 0.3048

// fmt.Printf("%g feet is %g metrs.\n", feet, meters)

// switch t := time.Now().Hour(); {

// case t > 6 && t < 12:
// 	fmt.Println("Good morning.")
// case t >= 12 && t < 18:
// 	fmt.Println("Good afternoon.")
// case t >= 18 && t < 23:
// 	fmt.Println("Good evening.")
// default:
// 	fmt.Println("Good night.")

// }

// switch h := time.Now().Hour(); {
// case h >= 18:
// 	fmt.Println("Good evening")
// case h >= 12:
// 	fmt.Println("Good afternoon")
// case h >= 6:
// 	fmt.Println("Good morning")
// default:
// 	fmt.Println("Good night")
// }

// Nested Loop Example

// }

// Nested Loop Example
// const max = 10

// func main() {
// 	fmt.Printf("%5s", "X")
// 	for i := 0; i <= max; i++ {
// 		fmt.Printf("%5d", i)
// 	}
// 	fmt.Println()
// 	for i := 0; i <= max; i++ {
// 		fmt.Printf("%5d", i)
// 		for j := 0; j <= max; j++ {
// 			fmt.Printf("%5d", i*j)
// 		}
// 		fmt.Println()
// 	}

// }

// Nested Loop Example 2

func main() {
	//words := strings.Fields("The crazy dog dance on the line")

	// for i := 0; i < len(words); i++ {
	// 	fmt.Printf("#%-2d: %q\n", i+1, words[i])
	// }
	// With Range Statement

	// for i, v := range words {
	// 	fmt.Printf("#%-2d: %q\n", i+1, v)
	// }

	// Also you can use blank-identifier for if you don't write the index variable
	// for _, v := range words {
	// 	fmt.Printf(" %q\n", v)
	// }

	// Random Usage

	rand.Seed(time.Now().UnixNano())

	//rand.Seed(100)
	guess := 10
	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()

}
