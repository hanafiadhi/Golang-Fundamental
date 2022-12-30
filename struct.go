package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var pelanggan Customer
	pelanggan.Name = "hanafi"
	pelanggan.Address = "cilacap"
	pelanggan.Age = 22

	fmt.Println(pelanggan.Name)

	joko := Customer{
		Name:    "Joko",
		Address: "Cilacap",
		Age:     23,
	}
	fmt.Println(joko)
}
