package main

import "fmt"

var (
	message string
)

func main() {
	message := "Hello World from Go !!"
	fmt.Println(message)
	var a [10]int
	fmt.Println("array a: ", a)
}
