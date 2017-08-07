package main

import "fmt"

// package level scope
// outside any containing block

var x int = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
