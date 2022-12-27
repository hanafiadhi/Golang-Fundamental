package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 70

	var luluscumload = ujian > 75

	var lulusAbsensi = absensi >= 70

	var lulus = lulusAbsensi && luluscumload

	fmt.Println(luluscumload)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)
}
