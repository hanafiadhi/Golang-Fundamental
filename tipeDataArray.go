package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Hanafi"
	name[1] = "adhi"
	name[2] = "prasetyo"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		80, 70, 90,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
