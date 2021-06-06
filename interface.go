package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	name := Person{
		Name: "World",
	}

	sayHello(name)

	cat := Animal{
		Name: "Yeaky",
	}

	sayHello(cat)
}
