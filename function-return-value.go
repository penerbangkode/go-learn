package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello World"
	} else {
		return "Hello " + name
	}
}
func main() {
	fmt.Println(getHello("Rido"))
	fmt.Println(getHello(""))
}
