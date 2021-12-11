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


## Abbreviations (Common abbrevations used in go)

The naming things such as variables, functions, paremeters is so important. There is a list for common abbrevations in go the following patterns.

```
var s string            // string
var i int               // index
var num int             // number
var msg string          // message
var v string            // value
var val string          // value
var fv string           // flag value
var err error           // error value
var args []string       // arguments
var seen bool           // has seen?
var parsed bool         // parsing ok?
var buff []byte         // buffer
var off int             // offset
var op int              // operation
var opRead              // read operation
var l int               // length
var n int               // number or number of
var m int               // another number
var c int               // capacity
var c int               // character
var r rune              // rune
var sep string          // separator
var src int             // source
var dst int             // destination
var b byte              // byte
var b []byte            // buffer
var buf []byte          // buffer
var w io.Writer         // writer
var r io.Reader         // reader
var pos int             // position

```
## Raw String Literal

 TYPES
 string   ->     "hi there" (string literal)
 string   ->	 `hi there` (raw string literal)

 Raw String Literal       |   String Literal
 -------------------------|---------------------
 multi-line		  |  single line
 not interpreted	  |  interpreted

 ```
 var s string
	s = "How are you?"
	s = `How are you?`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)

	s = `
	<html>
		<body>"Hello"</body>
	</html>`
	fmt.Println(s)
```
## String Length

Unicode characters can be 1 to 4 bytes each.

name := "john"
len(name) => 4
name := "İnanç"
len(name) => 7 ('İ' = 2 bytes, 'ç' = 2 bytes, the others are 1 bytes)

When you want to learn exact length, you can import "unicode/utf8" package. And you shoul write your code as below.

```
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "İnanç"

	fmt.Println(utf8.RuneCountInString(name))
}

```
Result should be 5. 
A rune can represent English and Non-English characters as well.



## TroubleShooting

a. When you have more than one go module in your workspace, the VS Code shows "gopls requires a module at the root of your workspace." error.
If that is the case, you can change go extension setting, in order to allow gopls to look for multiple modules in the workspace. To solve this, follow below steps.

	1 : Open Vscode, and then go to settings.

	2 : In the search bar , type gopls

	3 : Just below that you will find settings.json, click on that

	4 : Paste the below code their "gopls": { "experimentalWorkspaceModule": true, }

	5 : Save it and restart the Vscode, You are good to go now.
b. When you use os.Args and build your code on .exe file such as greeter.exe (Sample: go build -o greeter.exe) you can pay attention the following issue.
	When you change your code block in your file such as main.go, you should rebuild your code on .exe file for returning properly result.
