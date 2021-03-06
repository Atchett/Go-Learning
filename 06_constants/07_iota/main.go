package main

import "fmt"

// bit shifting - bitwise operations - binary
const (
	_  = iota             // 0 (ignore the first iota - 0)
	// move the bit 10 places to the left
	// you get the (in binary) 2 to the power 10 = 1024 (1Kb)
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	// move the bit 20 places to the left
	// you get the (in binary) 2 to the power 20 = 1,048,567 (1Mb)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", TB)
	fmt.Printf("%d\n", TB)
}
