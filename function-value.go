package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye" + name
}
func main() {
	goodbye := getGoodbye

	result := goodbye("hanafi")
	fmt.Println(result)
}
