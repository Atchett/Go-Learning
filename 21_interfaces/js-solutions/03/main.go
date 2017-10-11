package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	// sort forward and output
	sort.Ints(n)
	fmt.Println(n)
	// sort reverse and output
	// note: use of intSlice as this implements the sort.Interface interface
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
