package main

import "fmt"

// ordering of the variables matters
func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

// package level veriable so is accessible in any function within the package
// outer scope
var y = 42
