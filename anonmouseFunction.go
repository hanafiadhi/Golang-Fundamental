package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blaklist Blacklist) {
	if blaklist(name) {
		fmt.Println("You are blocked ", name)
	} else {
		fmt.Println("Wellcome aboard ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Hanafi", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
