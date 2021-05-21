package main

import "fmt"

func main() {
	var name = "Hello All"

	if name == "Hello" {
		fmt.Println("Hello")
	} else if name == "World" {
		fmt.Println("World")
	} else {
		fmt.Println("No")
	}

	// Shor statement
	if length := len(name); length > 5 {
		fmt.Println("Too Long!")
	} else {
		fmt.Println("Yeay, Correct")
	}
}
