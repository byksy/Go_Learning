package main

import (
	"fmt"
	"time"
)

func main() {

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

	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("Good evening")
	case h >= 12:
		fmt.Println("Good afternoon")
	case h >= 6:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good night")
	}

}
