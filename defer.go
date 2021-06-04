package main

import "fmt"

func logging() {
	fmt.Println("Done calling function")
}

func runApllication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)

}

func main() {
	runApllication(0)
}
