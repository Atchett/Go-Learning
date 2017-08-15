package main

import "fmt"

func zero(z int) {
	// z is a new variable that is being set to 0
	// passing the value not the variable
	z = 0
}

func main() {
	x := 5
	// passing 5 to x
	zero(x)
	// but then outputting the same number 5
	fmt.Println(x) // x is still 5
}
