package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Hanafi"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println(e)
	fmt.Println(eString)
}
