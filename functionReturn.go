package main

import "fmt"

func myName(name string) string {
	return "hello " + name
}

func main() {
	fmt.Println(myName("Hanafi"))
}
