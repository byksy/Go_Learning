package main

import (
	"fmt"
	"os"
	"strconv"
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

	arg := os.Args[1]

	feet, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("error: %q is not a number. \n", arg)
		return
	}
	meters := feet * 0.3048

	fmt.Printf("%g feet is %g metrs.\n", feet, meters)

}
