package main

import "fmt"

const (
	winter = 1
	summer = 3
	yearly = winter + summer
)

//test3
func main() {

	var books [yearly]string

	books[0] = "Kafka's Revenge"
	books[1] = "Stay Red"
	books[2] = "Fine Everything"
	books[3] = books[2] + " Let's Do It"

	fmt.Printf("books : %#v\n", books)

	var (
		wBooks [winter]string
		sBooks [summer]string
	)

	wBooks[0] = books[0]

	// IMPORTANT!!
	// range overed array gets copied into a new and hidden array variable.
	//Let's say the array contains millions of elements. They all would be copied.
	// for _, v /*This i a cop. It's not the original element*/ := range sBooks {
	// 	v += "won't effect"
	// }

	for i := range sBooks {
		sBooks[i] = books[i+1]
	}

	fmt.Printf("wBooks: %#v\n", wBooks)
	fmt.Printf("sBooks: %#v\n", sBooks)

	//If the length of an array belongs to compile time; then, why could I use the len function here?
	//Because the books' length is const value, when you compile your code the length of the array will be assigned.
	var published [len(books)]bool

	published[0] = true
	published[len(books)-1] = true

	fmt.Println("\nPublished Books: ")

	for i, ok := range published {
		if ok {
			fmt.Printf("+ %s\n", books[i])
		}
	}

}
