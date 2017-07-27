// Package stringutil contains utility functions for workign with strings
package stringutil

// Reverse returns it's argument string reversed rune-wise left to right
// https://blog.golang.org/strings
func Reverse(s string) string {
	return reverseTwo(s)
}

/*

go build
	go build reverse.go reverseTwo.go will not produce an output file.

go install
	will place the package inside the pkg directory of the workspace

*/