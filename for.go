package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	slice := []string{
		"Hanafi",
		"Adhi",
		"Prasetyo",
	}

	for _, value := range slice {
		fmt.Println(value)
	}
}
