package main

import "fmt"

func em(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return false
	}

}

func main() {
	var name interface{} = em(1)
	fmt.Println(name)
}
