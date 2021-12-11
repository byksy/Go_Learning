package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "İnanç"

	fmt.Println(utf8.RuneCountInString(name))
}
