package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var member1 Customer

	member1.Name = "Hello World"
	member1.Address = "Jakarta, Indonesia"
	member1.Age = 19

	fmt.Println(member1.Name)
	fmt.Println(member1.Address)
	fmt.Println(member1.Age)

	member2 := Customer{
		Name:    "Ella Musk",
		Address: "Bandung, Indonesia",
		Age:     22,
	}

	fmt.Println(member2)

	member3 := Customer{"Umeng", "Depok, Indonesia", 18}

	fmt.Println(member3)

}
