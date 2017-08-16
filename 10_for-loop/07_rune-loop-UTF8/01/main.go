package main

import "fmt"

func main() {
	for i := 50; i <= 140; i++ {
		// prints the int value,
		// the string representation of the value (i.e. the UTF-8 character
		// then convert to byte representation
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	// will give a
	foo := "a"
	// will give 97 (as it's a rune)
	bar := 'a'
	fmt.Println(foo)
	fmt.Println(bar)
	// will give the type of foo (which is a string)
	fmt.Printf("%T \n", foo)
	// will give the type of foo (which is a rune - or int32)
	fmt.Printf("%T \n", bar)
	// rune() is an alias to int32
	fmt.Printf("%T \n", rune(bar))

}
