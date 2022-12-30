package main

import "fmt"

type filter func(string) string

func sayHelloFilter(name string, filter filter) {
	filtername := filter(name)
	fmt.Println("Hello", filtername)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	sayHelloFilter("Hanafi", spamFilter)
	sayHelloFilter("anjing", spamFilter)
}
