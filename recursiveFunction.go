package main

import "fmt"

func recursive(nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return nilai * recursive(nilai-1)
	}
}

func main() {

	fmt.Println(recursive(5))
}
