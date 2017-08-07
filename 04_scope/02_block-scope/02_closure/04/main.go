package main

import "fmt"

// returns a function
func wrapper() func() int {
	// outer scope enclosing the inner scope
	x := 0
	return func() int {
		// inner scope
		x++
		return x
	}
}

func main() {
	// function expression
	increment := wrapper()
	// returns 1
	fmt.Println(increment())
	// returns 2
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
