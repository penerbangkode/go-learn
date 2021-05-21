package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 75
	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)
}
