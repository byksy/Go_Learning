package main

import (
	"fmt"
	"os"
)

//version:=0 DON'T DO
//var version string You can use that

func main() {

	fmt.Printf("%v#v\n", os.Args)

	fmt.Println("Path", os.Args[0])
	fmt.Println("1st arguument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])

	fmt.Println("Number of items inside os.Args:", len(os.Args))

}
