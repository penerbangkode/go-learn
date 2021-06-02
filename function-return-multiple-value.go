package main

import "fmt"

func getFullName() (string, string, string) {
	return "Hello", "&", "World"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
