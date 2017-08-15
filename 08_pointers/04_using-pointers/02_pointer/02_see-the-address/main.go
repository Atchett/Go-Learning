package main

import "fmt"

func zero(z *int) {
	// prints a memory address
	fmt.Println(z)
	*z = 0
}

func main() {
	x := 5
	// prints a memory address
	fmt.Println(&x)
	zero(&x)
	// prints a value
	fmt.Println(x) // x is 0
}
