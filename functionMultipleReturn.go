package main

import "fmt"

func getFullName() (string, string, string) {
	return "Hanafi", "Adhi", "Prasetyo"
}

func main() {
	firstname, middlename, _ := getFullName()

	fmt.Println(firstname)
	fmt.Println(middlename)
}
