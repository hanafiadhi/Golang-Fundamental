package main

import "fmt"

func main() {

	var month = [...]string{
		"Januari",
		"Febuari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"July",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = month[2:4]
	fmt.Println(slice2)
	var slice3 = append(slice2, "Hanafi")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Hanafi"
	fmt.Println(newSlice)
}
