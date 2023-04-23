package main

import "fmt"

// https://gobyexample.com/variables
func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable, only available inside a function
	f := "apple"
	fmt.Println(f)
}
