package main

import "fmt"

func main() {
	var name string

	name = "Hello World"
	fmt.Println(name)

	name = "Hello Mars"
	fmt.Println(name)

	var friendName = "Neptune"
	fmt.Println(friendName)

	var age = 10
	fmt.Println(age)

	ecosystem := "Milky Way"
	fmt.Println(ecosystem)

	ecosystem = "Sergeint"
	fmt.Println(ecosystem)

	var (
		firstName = "Hello"
		lastName  = "World"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
