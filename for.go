package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("No.", counter)
		counter++
	}

	// ini & post statment
	for number := 1; number <= 10; number++ {
		fmt.Println("Number =", number)
	}

	// array / slice loops

	slice := []string{"Hello", "World", "Sky", "Mars", "Moon"}

	for i := 1; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For range

	for i, data := range slice {
		fmt.Println(i, "=", data)
		// fmt.Println(data)

	}

	// Map

	detail := make(map[string]string)
	detail["world"] = "Earth"
	detail["country"] = "Indonesia"

	for key, data := range detail {
		fmt.Println(key, "=", data)
	}

}
