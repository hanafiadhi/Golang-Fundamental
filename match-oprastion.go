package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var (
		a = 10
		b = 10
		c = a * b
	)
	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)
}
