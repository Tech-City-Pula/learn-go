// Welcome to the challenges part of this repository.
// Please use the examples for reference and try to solve the challenges yourself.

// Make the tests succeed.
// run tests with: "go test -v"
package main

import (
	"fmt"
)

func HelloWorld() string {
	// HINT: make the helloString store the value "Hello, World!"
	const helloString = "Hello, !"
	fmt.Println(helloString)
	return helloString
}
