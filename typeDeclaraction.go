package main

import "fmt"

func main() {
	type NoKtp string

	var Ktp NoKtp = "033012391302"
	fmt.Println(Ktp)
}
