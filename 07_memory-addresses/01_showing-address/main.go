package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	// formatting the memory address as binary
	// note the Printf not Println
	fmt.Printf("%d \n", &a)
}
