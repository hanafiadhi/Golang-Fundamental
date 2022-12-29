package main

import "fmt"

func main() {
	var name = "adhi"
	if name == "hanafi" {
		fmt.Println(name)
	} else if name == "adhi" {
		fmt.Println(name)
	} else {
		fmt.Println("nama bukan Hanafi")
	}

	if lenght := len(name); lenght >= 4 {
		fmt.Println("panjang karakter ada 5")
	}
}
