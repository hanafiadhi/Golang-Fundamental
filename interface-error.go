package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai * pembagi, nil
	}
}
func main() {
	values, error := Pembagi(10, 5)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(values)
	}
}
