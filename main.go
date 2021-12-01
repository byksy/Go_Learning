package main

import (
	"fmt"
	"path"
)

//version:=0 DON'T DO
//var version string You can use that

func main() {

	//score := 0 DONT'
	//var score int //already score=0
	var dir, file string

	dir, file = path.Split("css/main.css")

	fmt.Println("file:", file)
	fmt.Println("dir:", dir)

	/*If you want discarding dir info you can use discard character*/
	// var file string
	// _ , file = path.Split("css/main.css")
	// fmt.Println ("file:",file)

	//// Grouping declaration
	// var (
	// 	//related
	// 	video string

	// 	//closely related
	// 	duration int
	// 	current int
	// )

}
