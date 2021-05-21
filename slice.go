package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Yes"

	// slice1[2] = "Test"
	// fmt.Println(slice1)

	var slice2 = months[9:11]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Tryna")
	slice3[1] = "Blue sky"

	fmt.Println(slice3)
	fmt.Println(months)

	var newSlice = make([]string, 2, 5)

	// Note:m
	// Copy berfungsi menduplikasi sebuah slice dan saat copyan diubah slice yang tercopy tidak ikut terubah. Sedangan jika membuat variable slice dan di isi dengan slice, slice yang di isikan tersebut ikut terubah jika variable tersebut diubah

	newSlice[0] = "arr 1"
	newSlice[1] = "arr 2"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	var testSlice []string = newSlice
	testSlice[1] = "slice test"
	fmt.Println(testSlice)

	copySlice[0] = "Copy change"
	fmt.Println(copySlice)
	fmt.Println(newSlice)

	// Note :
	// - Kalau ga ada capacity yang diberikan akan menjadi slice
	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
