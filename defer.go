package main

import "fmt"

func logging() {
	fmt.Println("memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Running aplication")
}
func main() {
	runApplication()
}
