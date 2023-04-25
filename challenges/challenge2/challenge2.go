// make the code compile and the tests succeed.
package main

import (
	"fmt"
)

// TODO: write a function that returns the sum of two integers
func Sum(a, b int) int {
	return ???
}

func main() {
	var sum1 = Sum(2, 3)
	fmt.Println(sum1)

	// Will this work?
	var sum2 = Sum(7, -2)
	fmt.Println(sum2)
}
