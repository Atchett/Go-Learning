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
	// sort reverse and output
	// note: use of stringSlice as this implements the sort.Interface interface
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
