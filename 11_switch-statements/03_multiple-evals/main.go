package main

import "fmt"

func main() {
	switch "Jenny" {
	// catches either Tim OR Jenny
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	// catches either Marcus OR Medhi
	case "Marcus", "Medhi":
		fmt.Println("Both of your names start with M")
	// catches either Julian OR Sushant
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}
