package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	hanafi := Man{"Hanafi"}
	hanafi.Married()
	fmt.Println(hanafi)
}