package main

import "fmt"

func main() {
	type text string
	type check bool

	var hello text = "Hello World"
	var selected check = true
	fmt.Println(hello)
	fmt.Println(selected)
}
