package main

import "fmt"

func random() interface{} {
	return "Hanafi"
}

func main() {
	var result interface{} = random()

	switch values := result.(type) {
	case string:
		fmt.Println("value", values, "is strinng")
	case int:
		fmt.Println("value", values, "is number")
	default:
		fmt.Println("unknow type")
	}
}
