package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Blue"
	names[1] = "In The"
	names[2] = "Sky"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10,
		11,
		12,
	}

	fmt.Println(values)

	var check [5]string
	fmt.Println(len(values))
	fmt.Println(len(check))
}
