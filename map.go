package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":    "Hanafi",
		"address": "Cilacap",
	}
	person["title"] = "Programer"

	fmt.Println(person["nama"])

	var book map[string]string = make(map[string]string)

	book["title"] = "Belajar Golang"
	book["author"] = "Hanafi"
	book["ups"] = "salah"
	fmt.Println(book)
	delete(book, "ups")
	fmt.Print(book)
}
