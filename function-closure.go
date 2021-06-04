package main

import "fmt"

func main() {
	say := "Hello"
	counter := 0

	increment := func() {
		say := "Hi"
		fmt.Println("Increment")
		counter++
		fmt.Println(say)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(say)
}
