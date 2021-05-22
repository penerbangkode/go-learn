package main

import "fmt"

func main() {
	name := "World"

	switch name {
	case "Hello":
		fmt.Println("Hello")
		fmt.Println("Hello")
	case "World":
		fmt.Println("World")
		fmt.Println("World")
	default:
		fmt.Println("Hello World")
		fmt.Println("Hello World")
	}

	// Short statement

	switch length := len(name); length >= 5 {
	case true:
		fmt.Println(true)
	case false:
		fmt.Println(false)
	}

	// Without expr
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("More than 10")
	case length > 5:
		fmt.Println("More than 5")
	default:
		fmt.Println("length =", length)
	}
}
