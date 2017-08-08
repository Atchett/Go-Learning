package main

import "fmt"

const (
	pi       = 3.14159
	language             = "Go"     // untyped
	typedLanguage string = "GoLang" // typed
)

func main() {
	fmt.Println(pi)
	fmt.Println(language)
	fmt.Println(typedLanguage)
}
