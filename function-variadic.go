package main

import "fmt"

func sumAll(numbers ...int) int {
	final := 0

	for _, value := range numbers {
		final += value
	}

	return final
}

func main() {
	final := sumAll(10, 15, 10, 10)
	fmt.Println(final)

	slice := []int{10, 20, 1, 5, 8}
	final = sumAll(slice...)
	fmt.Println(final)
}
