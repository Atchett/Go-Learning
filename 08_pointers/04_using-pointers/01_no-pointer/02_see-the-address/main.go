package main

import "fmt"

func zero(z int) {
	// the address of z (%p is base 16 (i.e. Hex) with the initial 0x added
	// note different to the address passed in
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func main() {
	// the address of x (%p is base 16 (i.e. Hex) with the initial 0x added
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
