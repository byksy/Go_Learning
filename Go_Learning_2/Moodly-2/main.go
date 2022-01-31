package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// args := os.Args[1:]
	// var n int
	// if len(args) != 2 {
	// 	fmt.Println("[yourname] [positive|negative]")
	// 	return
	// }

	// mood := [2][3]string{
	// 	{"feels ok 😍", "feels amazing 😊", "feels good 😉"},
	// 	{"feels bad 😒", "feels down 😌", "feels terrible 😪"},
	// }

	// rand.Seed(time.Now().UnixNano())

	// if args[1] == "positive" {

	// 	n = rand.Intn(len(mood[1]))

	// 	fmt.Printf("%s %s\n", args[0], mood[0][n])
	// } else {
	// 	n = rand.Intn(len(mood[0]))

	// 	fmt.Printf("%s %s\n", args[0], mood[1][n])
	// }

	/*REFACTORING*/

	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("[yourname] [positive|negative]")
		return
	}
	name, mood := args[0], args[1]

	moods := [2][3]string{
		{"ok 😍", "amazing 😊", "good 😉"},
		{"bad 😒", "down 😌", "terrible 😪"},
	}

	var mi int

	if mood == "negative" {
		mi = 1
	}

	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(len(moods[0]))

	fmt.Printf("\n%s feels %s\n", name, moods[mi][n])

}
