package main

import "fmt"

func main() {

	switch "Brian" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	case "Brian":
		fmt.Println("Wassup Brian")
		fallthrough
	case "Medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Wassup Julian")
	case "Sushant":
		fmt.Println("Wassup Sushant")
	}
	// will output Brian,
	// then falls through to Medhi,
	// and then falls through to Julian
	// then no one else (as no match)
}
