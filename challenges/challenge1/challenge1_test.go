package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var helloString = HelloWorld()
	if helloString != "Hello, World!" {
		t.Errorf("helloWorld() = %s; want Hello, World!", helloString)
	}
}
