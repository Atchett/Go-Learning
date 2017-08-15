package main

import "fmt"

// the * means it's a memory address storing an int
func zero(z *int) {
	// dereference and set a new value
	*z = 0
}

func main() {
	x := 5
	// passing the memory address
	zero(&x)
	fmt.Println(x) // x is 0
}
