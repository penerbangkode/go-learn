package main

import "fmt"

type Filter func(string) string

func sayHelloWithFillter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Noob" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFillter("World", spamFilter)
	sayHelloWithFillter("Noob", spamFilter)
}
