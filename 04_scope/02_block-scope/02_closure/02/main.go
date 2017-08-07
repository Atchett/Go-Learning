package main

import "fmt"

var x = 0

// returning an int
func increment() int {
	x++
	return x
}

func main() {
	// will return 1
	fmt.Println(increment())
	// will return 2
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
