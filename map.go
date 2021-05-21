package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Rido",
		"address": "address",
	}
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Rizervoir"
	book["author"] = "Me"
	book["status"] = "published"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "status")
	fmt.Println(book)
	fmt.Println(len(book))
}
