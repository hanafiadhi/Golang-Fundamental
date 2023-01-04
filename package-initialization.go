package main

import (
	"Belajar-Golang/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
