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
