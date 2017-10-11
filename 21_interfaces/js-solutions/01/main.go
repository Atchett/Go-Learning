package main

import (
	"fmt"
	"sort"
)

type people []string

// implementing the methods to implicityly implement the interface for Sort
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	// sort forward and output
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	// sort reverse and output
	// note: use of stringSlice as this implements the sort.Interface interface
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)
}
