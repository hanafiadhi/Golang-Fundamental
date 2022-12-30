package main

import "fmt"

type pelanggan struct {
	Name string
}

func (nama pelanggan) sayHello() {
	fmt.Println(nama.Name)
}

func main() {
	hanafi := pelanggan{Name: "hanafi"}
	hanafi.sayHello()
}
