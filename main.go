package main

import (
	"fmt"
	"path"
)

func main() {
	// fmt.Println(mascot.BestMascot())
	// fmt.Println(quote.Go())

	var dir, file string

	dir, file = path.Split("css/main.css")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
