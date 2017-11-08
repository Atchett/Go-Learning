package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	// sort forward and output
	sort.Strings(s)
	fmt.Println(s)
	//alternate method
	// sort.StringSlice implements the Interface interface on the s string slice so that we can now use .Sort
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	// another method
	// pass in the Interface interface which sort.StringSlice(s) implements
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
	// sort reverse and output
	// note: use of stringSlice as this implements the sort.Interface interface
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
