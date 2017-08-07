package main

import "fmt"

// lowercase only visible inside the package
// x only avaibale inside the function
// declare parameter
func max(x int) int {
	return 42 + x
}

func main() {
	// pass in argument
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
	// cannot now call the max function
}

// don't do this; bad coding practice to shadow variables
