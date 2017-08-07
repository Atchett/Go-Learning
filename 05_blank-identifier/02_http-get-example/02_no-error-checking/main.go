package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	// using a blank identifier (variable)
	// http.Get returns 2 variables, response and error
	res, _ := http.Get("http://www.bbc.co.uk")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
