# Go Learning
It is my learning program for Go language

## Short Declaration For Go
```
package main 
func main() {
    //var width, height =100,50 //DONT!
    width, height := 100,50 // it is ok!

    //DONT!
    //width =50 // assigns 50 to width
    // color := "red" // new variable: color

    width, color :=50, "red" //same as above, but it is so good!

}
```
## Type Conversion

```
	speed := 100 // speed is int
	force := 2.5 // for is float64

	//speed = speed * int(force) // You need to explicitly convert the values.
	// Conversion does not change the original value.
	// Return 200

	speed = int(float64(speed) * force) // Return 250. It is true.

	fmt.Println(speed)

```

## Basic os.Args

If you run your code block in Windows system. You can buil your application with following code.

```
go build -o greeter.exe
```
After that you can start your program with greeter.exe. Only write greeter.exe in to the VS Code Terminal and run it instead of go runn main.go 

For the below code block you can set three paremeters your .exe arguments. Then you can print your paremeters with the below code block.
Please don't forget the save your code before run your app.

```
   package main

import (
	"fmt"
	"os"
)


func main() {

	fmt.Printf("%v#v\n", os.Args)

	fmt.Println("Path", os.Args[0])
	fmt.Println("1st arguument:", os.Args[1])
	fmt.Println("2nd argument:", os.Args[2])
	fmt.Println("3rd argument:", os.Args[3])

	fmt.Println("Number of items inside os.Args:", len(os.Args))

}

```

After the writng code block you can run the following code block to see your result.

```
greeter.exe hi hello hey
```