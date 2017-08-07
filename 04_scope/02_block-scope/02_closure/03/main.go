package main

import "fmt"

func main() {
	// block level scoped var
	x := 0
	// anonymous function
	// assigning the function to increment
	// only way to define a function inside another function
	// a function expression
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
