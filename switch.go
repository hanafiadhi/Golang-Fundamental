package main

import "fmt"

func main() {
	var name = "adhi"
	switch length := len(name); length >= 4 {
	case true:
		fmt.Println("nama lebih besar sama dengan 4")
	case false:
		fmt.Println("nama kurang dari 4")
	}

	switch {
	case len(name) != 5:
		fmt.Println("ada")
	default:
		fmt.Println("nama kurang dari 5")
	}
}
