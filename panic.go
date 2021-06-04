package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error msg :", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App Error")
	}
	fmt.Println("App Success")
}

func main() {
	runApp(true)
	fmt.Println("Yeay!")
}
