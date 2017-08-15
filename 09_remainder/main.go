package main

import "fmt"

func main(){
	// % gives the remainder - similar to mod in JS
	x := 13 % 3
	fmt.Println(x)
	// the == is checking
	// = would assign the value
	if x == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 70; i++ {
		// note no () around the condition
		// can be in there if you want!
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
